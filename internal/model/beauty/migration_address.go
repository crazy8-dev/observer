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

package beauty

import (
	"context"

	"github.com/go-pg/pg/orm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/stats"

	"github.com/insolar/observer/internal/model"
)

type MigrationAddress struct {
	tableName struct{} `sql:"migration_addresses"`

	Addr      string `sql:",pk"`
	Timestamp int64  `sql:",notnull"`
	Wasted    bool
}

func (a *MigrationAddress) Dump(ctx context.Context, tx orm.DB) error {
	res, err := tx.Model(a).OnConflict("DO NOTHING").Insert(a)
	if err != nil {
		return errors.Wrapf(err, "failed to insert migration address")
	}

	if res.RowsAffected() == 0 {
		stats.Record(ctx, model.ErrorsCount.M(1))
		logrus.Errorf("Failed to insert migration address: %v", a)
	}
	return nil
}
