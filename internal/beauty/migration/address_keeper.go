//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package migration

import (
	"github.com/go-pg/pg"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	log "github.com/sirupsen/logrus"

	"github.com/insolar/observer/internal/dto"
	"github.com/insolar/observer/internal/model/beauty"
	"github.com/insolar/observer/internal/panic"
	"github.com/insolar/observer/internal/replication"
)

type Keeper struct {
	requests map[insolar.ID]*record.Material
	results  map[insolar.ID]*record.Material
	cache    []*beauty.WasteMigrationAddress

	migrationAddressGauge prometheus.Gauge
	stat                  *dumpStat
}

func NewKeeper(migrationAddressGauge prometheus.Gauge) *Keeper {
	stat := &dumpStat{
		cached: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "observer_migration_address_keeper_cached_total",
			Help: "Cache size of migration address composer",
		}),
	}

	return &Keeper{
		requests:              make(map[insolar.ID]*record.Material),
		results:               make(map[insolar.ID]*record.Material),
		migrationAddressGauge: migrationAddressGauge,
		stat:                  stat,
	}
}

func (k *Keeper) Process(rec *record.Material) {
	defer panic.Log("migration_address_keeper")

	switch v := rec.Virtual.Union.(type) {
	case *record.Virtual_Result:
		origin := *v.Result.Request.GetLocal()
		if req, ok := k.requests[origin]; ok {
			delete(k.requests, origin)
			if isGetFreeMigrationAddress(req) {
				k.processResult(rec)
			}
		} else {
			k.results[origin] = rec
		}
	case *record.Virtual_IncomingRequest:
		origin := rec.ID
		if res, ok := k.results[origin]; ok {
			delete(k.results, origin)
			if isGetFreeMigrationAddress(rec) {
				k.processResult(res)
			}
		} else {
			k.requests[origin] = rec
		}
	case *record.Virtual_OutgoingRequest:
		origin := rec.ID
		if _, ok := k.results[origin]; ok {
			delete(k.results, origin)
		} else {
			k.requests[origin] = rec
		}
	}
}

func (k *Keeper) Dump(tx *pg.Tx, pub replication.OnDumpSuccess) error {
	log.Infof("dump wasted migration addresses")
	k.updateStat()

	deferred := []*beauty.WasteMigrationAddress{}
	for _, addr := range k.cache {
		if err := addr.Dump(tx); err != nil {
			deferred = append(deferred, addr)
		}
	}

	for _, addr := range deferred {
		log.Infof("Deferred migration address %s", addr.Addr)
	}

	pub.Subscribe(func() {
		subtrahend := len(k.cache) - len(deferred)
		k.migrationAddressGauge.Sub(float64(subtrahend))
		k.cache = deferred
	})
	return nil
}

func (k *Keeper) processResult(rec *record.Material) {
	result := (*dto.Result)(rec)
	if !result.IsSuccess() {
		return
	}
	addr := wastedAddress(result)

	k.cache = append(k.cache, &beauty.WasteMigrationAddress{
		Addr: addr,
	})
}

func isGetFreeMigrationAddress(rec *record.Material) bool {
	v, ok := rec.Virtual.Union.(*record.Virtual_IncomingRequest)
	if !ok {
		return false
	}

	in := v.IncomingRequest
	return in.Method == "GetFreeMigrationAddress"
}

func wastedAddress(result *dto.Result) string {
	rets := result.ParsePayload().Returns
	address, ok := rets[0].(string)
	if !ok {
		return ""
	}
	return address
}

func (k *Keeper) updateStat() {
	requestCount := len(k.requests)
	resultCount := len(k.results)

	k.stat.cached.Set(float64(requestCount + resultCount))
}
