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
	"time"

	"github.com/insolar/insolar/insolar"

	"github.com/insolar/observer/configuration"
	"github.com/insolar/observer/connectivity"
	"github.com/insolar/observer/internal/app/observer"
	"github.com/insolar/observer/internal/pkg/panic"
	"github.com/insolar/observer/observability"
)

type Manager struct {
	stop chan bool

	cfg      *configuration.Configuration
	init     func() *state
	fetch    func(*state) *raw
	beautify func(*raw) *beauty
	filter   func(*beauty) *beauty
	store    func(*beauty, *state) *observer.Statistic

	router *Router
}

func Prepare() *Manager {
	cfg := configuration.Load()
	obs := observability.Make()
	conn := connectivity.Make(cfg, obs)
	return &Manager{
		stop:     make(chan bool, 1),
		init:     makeInitter(cfg, obs, conn),
		fetch:    makeFetcher(cfg, obs, conn),
		beautify: makeBeautifier(obs),
		filter:   makeFilter(obs),
		store:    makeStorer(cfg, obs, conn),
		router:   NewRouter(cfg, obs),
		cfg:      cfg,
	}
}

func (m *Manager) Start() {
	go func() {
		defer panic.Catch("component.Manager")

		m.router.Start()
		defer m.router.Stop()

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
	m.stop <- true
}

func (m *Manager) needStop() bool {
	select {
	case <-m.stop:
		return true
	default:
		// continue
	}
	return false
}

func (m *Manager) run(s *state) {
	raw := m.fetch(s)
	beauty := m.beautify(raw)
	collapsed := m.filter(beauty)
	statistic := m.store(collapsed, s)

	if raw != nil {
		s.last = raw.pulse.Number
	}

	if statistic != nil {
		s.stat = *statistic
	}

	time.Sleep(m.cfg.Replicator.AttemptInterval)
}

type raw struct {
	pulse             *observer.Pulse
	batch             []*observer.Record
	shouldIterateFrom insolar.PulseNumber
}

type beauty struct {
	pulse       *observer.Pulse
	records     []*observer.Record
	requests    []*observer.Request
	results     []*observer.Result
	activates   []*observer.Activate
	amends      []*observer.Amend
	deactivates []*observer.Deactivate

	transfers []*observer.DepositTransfer
	members   map[insolar.ID]*observer.Member
	balances  map[insolar.ID]*observer.Balance
	deposits  map[insolar.ID]*observer.Deposit
	updates   map[insolar.ID]*observer.DepositUpdate
	addresses map[string]*observer.MigrationAddress
	wastings  map[string]*observer.Wasting
}

type state struct {
	last insolar.PulseNumber
	rp   RecordPosition
	stat observer.Statistic
}

type RecordPosition struct {
	Last              insolar.PulseNumber
	RN                uint32
	ShouldIterateFrom insolar.PulseNumber
}
