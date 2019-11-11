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

package component

import (
	"context"
	"time"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/ledger/heavy/exporter"
	"github.com/sirupsen/logrus"

	"github.com/insolar/observer/configuration"
	"github.com/insolar/observer/connectivity"
	"github.com/insolar/observer/internal/app/observer"
	"github.com/insolar/observer/internal/app/observer/grpc"
	"github.com/insolar/observer/observability"
)

type Manager struct {
	stopSignal chan bool

	cfg           *configuration.Configuration
	log           logrus.Logger
	init          func() *state
	commonMetrics *observability.CommonObserverMetrics
	fetch         func(context.Context, *state) *raw
	beautify      func(context.Context, *raw) *beauty
	filter        func(*beauty) *beauty
	store         func(*beauty, *state) *observer.Statistic
	stop          func()

	router       RouterInterface
	sleepCounter sleepCounter
}

func Prepare() *Manager {
	cfg := configuration.Load()
	obs := observability.Make(cfg)
	conn := connectivity.Make(cfg, obs)
	router := NewRouter(cfg, obs)
	pulses := grpc.NewPulseFetcher(cfg, obs, exporter.NewPulseExporterClient(conn.GRPC()))
	records := grpc.NewRecordFetcher(cfg, obs, exporter.NewRecordExporterClient(conn.GRPC()))
	sm := NewSleepManager(cfg)
	return &Manager{
		stopSignal:    make(chan bool, 1),
		init:          makeInitter(cfg, obs, conn),
		log:           *obs.Log(),
		commonMetrics: observability.MakeCommonMetrics(obs),
		fetch:         makeFetcher(obs, pulses, records),
		beautify:      makeBeautifier(cfg, obs, conn),
		filter:        makeFilter(obs),
		store:         makeStorer(cfg, obs, conn),
		stop:          makeStopper(obs, conn, router),
		router:        router,
		cfg:           cfg,
		sleepCounter:  sm,
	}
}

func (m *Manager) Start() {
	go func() {
		m.router.Start()
		defer m.stop()

		state := m.init()
		for {
			m.run(state)
			if m.needStop() {
				return
			}
		}
	}()
}

func (m *Manager) Stop() {
	m.stopSignal <- true
}

func (m *Manager) needStop() bool {
	select {
	case <-m.stopSignal:
		return true
	default:
		// continue
	}
	return false
}

func (m *Manager) run(s *state) {
	timeStart := time.Now()
	ctx := context.Background()
	raw := m.fetch(ctx, s)
	beauty := m.beautify(ctx, raw)
	collapsed := m.filter(beauty)
	statistic := m.store(collapsed, s)

	timeExecuted := time.Since(timeStart)
	m.commonMetrics.PulseProcessingTime.Set(timeExecuted.Seconds())
	m.log.Debug("timeExecuted: ", timeExecuted)
	m.log.Debugf("Stats: %+v", statistic)

	if raw != nil {
		s.last = raw.pulse.Number
		s.rp.ShouldIterateFrom = raw.shouldIterateFrom
	}

	sleepTime := m.sleepCounter.Count(ctx, raw, timeExecuted)
	m.log.Debug("Sleep: ", sleepTime)
	time.Sleep(sleepTime)
}

type raw struct {
	pulse             *observer.Pulse
	batch             map[uint32]*exporter.Record
	shouldIterateFrom insolar.PulseNumber
	currentHeavyPN    insolar.PulseNumber
}

type beauty struct {
	pulse       *observer.Pulse
	records     map[uint32]*exporter.Record
	requests    []*observer.Request
	results     []*observer.Result
	activates   []*observer.Activate
	amends      []*observer.Amend
	deactivates []*observer.Deactivate

	members        map[insolar.ID]*observer.Member
	balances       map[insolar.ID]*observer.Balance
	deposits       map[insolar.ID]*observer.Deposit
	depositUpdates map[insolar.ID]*observer.DepositUpdate
	addresses      map[string]*observer.MigrationAddress
	wastings       map[string]*observer.Wasting

	txRegister   []observer.TxRegister
	txResult     []observer.TxResult
	txSagaResult []observer.TxSagaResult
}

type state struct {
	last insolar.PulseNumber
	rp   RecordPosition
	ms   metricState
}

type RecordPosition struct {
	Last              insolar.PulseNumber
	RN                uint32
	ShouldIterateFrom insolar.PulseNumber
}
