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
	"github.com/go-pg/pg/orm"
	"github.com/pkg/errors"

	"github.com/insolar/observer/v2/configuration"
	"github.com/insolar/observer/v2/connectivity"
	"github.com/insolar/observer/v2/internal/app/observer/postgres"
	"github.com/insolar/observer/v2/observability"
)

func initDB(cfg *configuration.Configuration, obs *observability.Observability, conn *connectivity.Connectivity) {
	log := obs.Log()
	if cfg == nil {
		return
	}
	if cfg.DB.CreateTables {
		db := conn.PG()

		err := db.CreateTable(&postgres.TransferSchema{}, &orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			log.Error(errors.Wrapf(err, "failed to create transfers table"))
		}
	}
}
