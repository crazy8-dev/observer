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

package models

import (
	"fmt"
	"github.com/insolar/insolar/pulse"
	"reflect"
	"sync"
)

type Member struct {
	tableName struct{} `sql:"members"` //nolint: unused,structcheck

	Reference        []byte `sql:"member_ref"`
	WalletReference  []byte `sql:"wallet_ref"`
	AccountReference []byte `sql:"account_ref"`
	AccountState     []byte `sql:"account_state"`
	MigrationAddress string `sql:"migration_address"`
	Balance          string `sql:"balance"`
	Status           string `sql:"status"`
}

type Deposit struct {
	tableName struct{} `sql:"deposits"` //nolint: unused,structcheck

	Reference       []byte `sql:"deposit_ref"`
	MemberReference []byte `sql:"member_ref"`
	EtheriumHash    string `sql:"eth_hash"`
	State           []byte `sql:"deposit_state"`
	HoldReleaseDate int64  `sql:"hold_release_date"`
	Amount          string `sql:"varchar"`
	Balance         string `sql:"balance"`
}

type TransactionStatus string

const (
	TStatusUnknown  TransactionStatus = "unknown"
	TStatusPending  TransactionStatus = "pending"
	TStatusSent     TransactionStatus = "sent"
	TStatusReceived TransactionStatus = "received"
	TStatusFailed   TransactionStatus = "failed"
)

type TransactionType string

const (
	TTypeUnknown   TransactionType = "unknown"
	TTypeTransfer  TransactionType = "transfer"
	TTypeMigration TransactionType = "migration"
	TTypeRelease   TransactionType = "release"
)

type Transaction struct {
	tableName struct{} `sql:"simple_transactions"` //nolint: unused,structcheck

	// Indexes.
	ID            int64  `sql:"id"`
	TransactionID []byte `sql:"tx_id"`

	// Request registered.
	StatusRegistered     bool            `sql:"status_registered"`
	Type                 TransactionType `sql:"type"`
	PulseRecord          [2]int64        `sql:"pulse_record" pg:",array"`
	MemberFromReference  []byte          `sql:"member_from_ref"`
	MemberToReference    []byte          `sql:"member_to_ref"`
	DepositToReference   []byte          `sql:"deposit_to_ref"`
	DepositFromReference []byte          `sql:"deposit_from_ref"`
	Amount               string          `sql:"amount"`
	Fee                  string          `sql:"fee"`

	// Result received.
	StatusSent bool `sql:"status_sent"`

	// Saga result received.
	StatusFinished    bool     `sql:"status_finished"`
	FinishSuccess     bool     `sql:"finish_success"`
	FinishPulseRecord [2]int64 `sql:"finish_pulse_record" pg:",array"`
}

type fieldCache struct {
	sync.Mutex
	cache map[reflect.Type][]string
}

var fieldsCache = fieldCache{
	cache: make(map[reflect.Type][]string),
}

func (t Transaction) Fields() []string {
	fieldsCache.Lock()
	defer fieldsCache.Unlock()

	tType := reflect.TypeOf(t)

	if fields, ok := fieldsCache.cache[tType]; ok {
		return append(fields[:0:0], fields...)
	}
	fieldsCache.cache[tType] = getFieldList(tType)
	fields := fieldsCache.cache[tType]
	return append(fields[:0:0], fields...)
}

func (t Transaction) QuotedFields() []string {
	fields := t.Fields()
	for i := range fields {
		fields[i] = fmt.Sprintf("'%s'", fields[i])
	}
	return fields
}

func getFieldList(t reflect.Type) []string {
	var fieldList []string

	for i := 0; i < t.NumField(); i++ {
		// ignore tableName
		if t.Field(i).Name == "tableName" {
			continue
		}
		tag := t.Field(i).Tag.Get("sql")
		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		fieldList = append(fieldList, tag)
	}

	return fieldList
}

func (t *Transaction) Status() TransactionStatus {
	registered := t.StatusRegistered
	sent := t.StatusRegistered && t.StatusSent
	finished := t.StatusRegistered && t.StatusFinished

	if finished {
		if t.FinishSuccess {
			return TStatusReceived
		}
		return TStatusFailed
	}
	if sent {
		return TStatusSent
	}
	if registered {
		return TStatusPending
	}

	return TStatusUnknown
}

func (t *Transaction) PulseNumber() int64 {
	return t.PulseRecord[0]
}

func (t *Transaction) RecordNumber() int64 {
	return t.PulseRecord[1]
}

func (t *Transaction) Index() string {
	return fmt.Sprintf("%d:%d", t.PulseRecord[0], t.PulseRecord[1])
}

func (t *Transaction) Timestamp() float32 {
	p := t.PulseNumber()
	pulseTime, err := pulse.Number(p).AsApproximateTime()
	if err != nil {
		return 0
	}
	return float32(pulseTime.Unix())
}
