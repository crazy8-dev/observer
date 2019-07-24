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

package sequence

import (
	"bytes"
	"sync"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/ledger/heavy/sequence"
	"github.com/pkg/errors"

	"github.com/insolar/observer/internal/ledger/store"
)

type sequencer struct {
	sync.RWMutex
	db store.DB
}

func NewMimicSequencer(db store.DB) sequence.Sequencer {
	return &sequencer{db: db}
}

func (s *sequencer) Len(scope byte, pulse insolar.PulseNumber) uint32 {
	s.RLock()
	defer s.RUnlock()

	result := 0
	pivot := polyKey{id: []byte{}, scope: store.Scope(scope)}
	it := s.db.NewIterator(pivot, false)
	defer it.Close()
	for it.Next() && bytes.HasPrefix(it.Key(), pulse.Bytes()) {
		result++
	}
	return uint32(result)
}

func (s *sequencer) First(scope byte) *sequence.Item {
	pivot := polyKey{id: []byte{}, scope: store.Scope(scope)}
	it := s.db.NewIterator(pivot, false)
	defer it.Close()

	if !it.Next() {
		return nil
	}
	return &sequence.Item{Key: it.Key(), Value: it.Value()}
}

func (s *sequencer) Last(scope byte) *sequence.Item {
	pivot := polyKey{id: []byte{0xFF, 0xFF, 0xFF, 0xFF}, scope: store.Scope(scope)}
	it := s.db.NewIterator(pivot, true)
	defer it.Close()

	if !it.Next() {
		return nil
	}
	return &sequence.Item{Key: it.Key(), Value: it.Value()}
}

func (s *sequencer) Slice(scope byte, from insolar.PulseNumber, skip uint32, limit uint32) []sequence.Item {
	s.RLock()
	defer s.RUnlock()

	var result []sequence.Item
	pivot := polyKey{id: from.Bytes(), scope: store.Scope(scope)}
	it := s.db.NewIterator(pivot, false)
	defer it.Close()

	skipped := 0
	for it.Next() && pulse(it.Key()) == from && len(result) < int(limit) {
		if skipped < int(skip) {
			skipped++
			continue
		}
		result = append(result, sequence.Item{
			Key:   it.Key(),
			Value: it.Value(),
		})
	}
	return result
}

func (s *sequencer) Upsert(scope byte, sequence []sequence.Item) error {
	s.Lock()
	defer s.Unlock()

	for _, item := range sequence {
		key := polyKey{id: item.Key, scope: store.Scope(scope)}
		err := s.db.Set(key, item.Value)
		if err != nil {
			return errors.Wrapf(err, "failed to save item of sequence")
		}
	}
	return nil
}

type polyKey struct {
	id    []byte
	scope store.Scope
}

func (k polyKey) ID() []byte {
	return k.id
}

func (k polyKey) Scope() store.Scope {
	return k.scope
}

func pulse(buf []byte) insolar.PulseNumber {
	return insolar.NewPulseNumber(buf)
}
