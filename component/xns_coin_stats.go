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
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/insolar/observer/internal/app/observer/postgres"
)

type XnsCoinStats struct {
	Created     time.Time `json:"-"`
	Total       string    `json:"total"`
	Max         string    `json:"max"`
	Circulating string    `json:"circulating"`
}

type StatsGetter interface {
	Coins() (XnsCoinStats, error)
	Total() (string, error)
	Max() (string, error)
	Circulating() (string, error)
}

type StatsManager struct {
	log        *logrus.Logger
	repository postgres.StatsRepo
}

func NewStatsManager(log *logrus.Logger, r postgres.StatsRepo) *StatsManager {
	return &StatsManager{
		log:        log,
		repository: r,
	}
}

func (s *StatsManager) Coins() (XnsCoinStats, error) {
	lastStats, err := s.repository.LastStats()
	if err != nil {
		return XnsCoinStats{}, errors.Wrap(err, "failed request get stats")
	}
	return XnsCoinStats{
		Created:     lastStats.Created,
		Total:       s.convertToCMCFormat(lastStats.Total),
		Max:         s.convertToCMCFormat(lastStats.Max),
		Circulating: s.convertToCMCFormat(lastStats.Circulating),
	}, nil
}

func (s *StatsManager) Total() (string, error) {
	res, err := s.Coins()
	if err != nil {
		return "", err
	}
	return res.Total, nil
}

func (s *StatsManager) Max() (string, error) {
	res, err := s.Coins()
	if err != nil {
		return "", err
	}
	return res.Max, nil
}

func (s *StatsManager) Circulating() (string, error) {
	res, err := s.Coins()
	if err != nil {
		return "", err
	}
	return res.Circulating, nil
}

func (s *StatsManager) CountStats() (XnsCoinStats, error) {
	st, err := s.repository.CountStats()
	if err != nil {
		return XnsCoinStats{}, err
	}
	return s.toDTO(st), nil
}

func (s *StatsManager) InsertStats(xcs XnsCoinStats) error {
	return s.repository.InsertStats(s.fromDTO(xcs))
}

func (s *StatsManager) toDTO(stats postgres.StatsModel) XnsCoinStats {
	return XnsCoinStats{
		Created:     stats.Created,
		Total:       stats.Total,
		Max:         stats.Max,
		Circulating: stats.Circulating,
	}
}

func (s *StatsManager) fromDTO(stats XnsCoinStats) postgres.StatsModel {
	return postgres.StatsModel{
		Created:     stats.Created,
		Total:       stats.Total,
		Max:         stats.Max,
		Circulating: stats.Circulating,
	}
}

func (s *StatsManager) convertToCMCFormat(str string) string {
	if len(str) <= 10 {
		return fmt.Sprintf("0.%010s", str)
	}
	return str[:len(str)-10] + "." + str[len(str)-10:]
}
