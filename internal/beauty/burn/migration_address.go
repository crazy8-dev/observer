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

package burn

import (
	"encoding/json"

	"github.com/go-pg/pg"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/logicrunner/builtin/contract/member"
	"github.com/insolar/insolar/logicrunner/builtin/contract/member/signer"
	"github.com/insolar/insolar/network/consensus/common/pulse"
	"github.com/pkg/errors"

	"github.com/insolar/observer/internal/model/beauty"
	"github.com/insolar/observer/internal/replication"
	log "github.com/sirupsen/logrus"
)

type Composer struct {
	requests map[insolar.ID]*record.Material
	results  map[insolar.ID]*record.Material
	cache    []*beauty.MigrationAddress
}

func NewComposer() *Composer {
	return &Composer{
		requests: make(map[insolar.ID]*record.Material),
		results:  make(map[insolar.ID]*record.Material),
	}
}

func (c *Composer) Process(rec *record.Material) {
	switch v := rec.Virtual.Union.(type) {
	case *record.Virtual_Result:
		origin := *v.Result.Request.Record()
		if req, ok := c.requests[origin]; ok {
			delete(c.requests, origin)
			if isAddBurnAddresses(req) {
				c.processAddBurnAddresses(req, rec)
			}
		} else {
			c.results[origin] = rec
		}
	case *record.Virtual_IncomingRequest:
		origin := rec.ID
		if res, ok := c.results[origin]; ok {
			delete(c.results, origin)
			if isAddBurnAddresses(rec) {
				c.processAddBurnAddresses(rec, res)
			}
		} else {
			c.requests[origin] = rec
		}
	case *record.Virtual_OutgoingRequest:
		origin := rec.ID
		if _, ok := c.results[origin]; ok {
			delete(c.results, origin)
		} else {
			c.requests[origin] = rec
		}
	}
}

func (c *Composer) Dump(tx *pg.Tx, pub replication.OnDumpSuccess) error {
	for _, addr := range c.cache {
		if err := addr.Dump(tx); err != nil {
			return errors.Wrapf(err, "failed to dump migration addresses addr=%s", addr.Addr)
		}
	}

	pub.Subscribe(func() {
		c.cache = []*beauty.MigrationAddress{}
	})
	return nil
}

func (c *Composer) processAddBurnAddresses(req *record.Material, res *record.Material) {
	if !successResult(res) {
		return
	}
	in := req.Virtual.GetIncomingRequest()
	args := parseCallArguments(in.Arguments)
	addresses := parseAddBurnAddressesCallParams(args)
	pn := pulse.Number(req.ID.Pulse())
	for _, addr := range addresses {
		c.cache = append(c.cache, &beauty.MigrationAddress{
			Addr:      addr,
			Timestamp: pn.AsApproximateTime().Unix(),
			Wasted:    false,
		})
	}
}

func isAddBurnAddresses(rec *record.Material) bool {
	v, ok := rec.Virtual.Union.(*record.Virtual_IncomingRequest)
	if !ok {
		return false
	}

	in := v.IncomingRequest
	if in.Method != "Call" {
		return false
	}

	args := parseCallArguments(in.Arguments)
	return args.Params.CallSite == "migration.addBurnAddresses"
}

func parseCallArguments(inArgs []byte) member.Request {
	var args []interface{}
	err := insolar.Deserialize(inArgs, &args)
	if err != nil {
		log.Warn(errors.Wrapf(err, "failed to deserialize request arguments"))
		return member.Request{}
	}

	request := member.Request{}
	if len(args) > 0 {
		if rawRequest, ok := args[0].([]byte); ok {
			var (
				pulseTimeStamp int64
				signature      string
				raw            []byte
			)
			err = signer.UnmarshalParams(rawRequest, &raw, &signature, &pulseTimeStamp)
			if err != nil {
				log.Warn(errors.Wrapf(err, "failed to unmarshal params"))
				return member.Request{}
			}
			err = json.Unmarshal(raw, &request)
			if err != nil {
				log.Warn(errors.Wrapf(err, "failed to unmarshal json member request"))
				return member.Request{}
			}
		}
	}
	return request
}

func successResult(res *record.Material) bool {
	payload := res.Virtual.GetResult().Payload
	params := parsePayload(payload)
	if len(params) < 2 {
		log.Error(errors.New("failed to parse migration.addBurnAddresses result payload"))
		return false
	}
	ret1 := params[1]
	if ret1 != nil {
		errMap, ok := ret1.(map[string]interface{})
		if !ok {
			log.Error(errors.New("failed to parse migration.addBurnAddresses result payload"))
			return false
		}

		strRepresentation, ok := errMap["S"]
		if !ok {
			log.Error(errors.New("failed to parse migration.addBurnAddresses result payload"))
			return false
		}

		msg, ok := strRepresentation.(string)
		if !ok {
			log.Error(errors.New("failed to parse migration.addBurnAddresses result payload"))
			return false
		}

		log.Error(errors.New(msg))
		return false
	}

	return true
}