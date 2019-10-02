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

	"github.com/insolar/observer/v2/configuration"
	"github.com/insolar/observer/v2/connectivity"
	"github.com/insolar/observer/v2/internal/app/observer"
	"github.com/insolar/observer/v2/internal/pkg/panic"
	"github.com/insolar/observer/v2/observability"
)

type Manager struct {
	stop chan bool

	cfg      *configuration.Configuration
	init     func() *state
	fetch    func(*state) *raw
	beautify func(*raw) *beauty
	store    func(*beauty)

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
		store:    makeStorer(obs, conn),
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
	m.store(beauty)

	if raw != nil {
		s.last = raw.pulse.Number
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
	members   []*observer.Member
	balances  []*observer.Balance
	deposits  []*observer.Deposit
	updates   []*observer.DepositUpdate
	addresses []*observer.MigrationAddress
	wastings  []*observer.Wasting
	users     []*observer.User
	groups    []*observer.Group
}

type state struct {
	last insolar.PulseNumber
	rp   RecordPosition
}

type RecordPosition struct {
	Last              insolar.PulseNumber
	RN                uint32
	ShouldIterateFrom insolar.PulseNumber
}
