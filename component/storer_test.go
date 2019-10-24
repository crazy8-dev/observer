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
// +build integration

package component

import (
	"bytes"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/insolar/observer/configuration"
	"github.com/insolar/observer/connectivity"
	"github.com/insolar/observer/internal/app/observer"
	"github.com/insolar/observer/internal/models"
	"github.com/insolar/observer/observability"
)

func Test_makeStorer(t *testing.T) {
	cfg := configuration.Default()
	obs := observability.Make(cfg)
	conn := connectivity.Make(cfg, obs)
	storer := makeStorer(cfg, obs, conn)

	b := &beauty{
		transfers: []*observer.Transfer{{}},
	}
	s := &state{}

	cfg.DB.Attempts = 1
	cfg.DB.AttemptInterval = time.Nanosecond

	require.NotPanics(t, func() {
		storer(b, s)
	})
}

func TestStoreSimpleTransactions(t *testing.T) {
	expectedTransactions := []models.Transaction{
		{
			TransactionID:       []byte{byte(rand.Int())},
			PulseNumber:         rand.Int63(),
			MemberFromReference: []byte{byte(rand.Int())},
			MemberToReference:   []byte{byte(rand.Int())},
			Amount:              strconv.Itoa(rand.Int()),
			Fee:                 strconv.Itoa(rand.Int()),
			FinishSuccess:       rand.Int()/2 == 0,
			FinishPulseNumber:   rand.Int63(),
			FinishRecordNumber:  rand.Int63(),
			StatusRegistered:    true,
			StatusSent:          true,
			StatusFinished:      true,
		},
		{
			TransactionID:         []byte{byte(rand.Int())},
			PulseNumber:           rand.Int63(),
			MigrationsToReference: []byte{byte(rand.Int())},
			Amount:                strconv.Itoa(rand.Int()),
			Fee:                   strconv.Itoa(rand.Int()),
			StatusRegistered:      true,
			StatusSent:            true,
			StatusFinished:        false,
		},
	}
	_ = db.RunInTransaction(func(tx *pg.Tx) error {
		// Create different update functions.
		funcs := []func() error{
			func() error {
				return StoreTxRegister(tx, []observer.TxRegister{
					{
						TransactionID:       expectedTransactions[0].TransactionID,
						PulseNumber:         expectedTransactions[0].PulseNumber,
						MemberFromReference: expectedTransactions[0].MemberFromReference,
						MemberToReference:   expectedTransactions[0].MemberToReference,
						Amount:              expectedTransactions[0].Amount,
						Fee:                 expectedTransactions[0].Fee,
					},
					{
						TransactionID:         expectedTransactions[1].TransactionID,
						PulseNumber:           expectedTransactions[1].PulseNumber,
						MigrationsToReference: expectedTransactions[1].MigrationsToReference,
						Amount:                expectedTransactions[1].Amount,
						Fee:                   expectedTransactions[1].Fee,
					},
				})
			},
			func() error {
				return StoreTxResult(tx, []observer.TxResult{
					{
						TransactionID: expectedTransactions[0].TransactionID,
					},
					{
						TransactionID: expectedTransactions[1].TransactionID,
					},
				})
			},
			func() error {
				return StoreTxSagaResult(tx, []observer.TxSagaResult{
					{
						TransactionID:      expectedTransactions[0].TransactionID,
						FinishSuccess:      expectedTransactions[0].FinishSuccess,
						FinishPulseNumber:  expectedTransactions[0].FinishPulseNumber,
						FinishRecordNumber: expectedTransactions[0].FinishRecordNumber,
					},
				})
			},
		}

		// Run functions in random order.
		rand.Shuffle(len(funcs), func(i, j int) {
			t := funcs[i]
			funcs[i] = funcs[j]
			funcs[j] = t
		})
		for _, f := range funcs {
			err := f()
			require.NoError(t, err)
		}

		// Select transactions from db.
		selected := make([]models.Transaction, 2)
		res, err := tx.Query(&selected, `SELECT * FROM simple_transactions ORDER BY tx_id`)
		require.NoError(t, err)
		require.Equal(t, 2, res.RowsReturned())
		// Reset ID field to simplify comparing.
		for i, t := range selected {
			t.ID = 0
			selected[i] = t
		}
		// Sort expected slice.
		sort.Slice(expectedTransactions, func(i, j int) bool {
			return bytes.Compare(expectedTransactions[i].TransactionID, expectedTransactions[j].TransactionID) == -1
		})
		// Compare transactions.
		assert.Equal(t, expectedTransactions, selected)
		return tx.Rollback()
	})
}
