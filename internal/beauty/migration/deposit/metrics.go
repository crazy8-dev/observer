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

package deposit

import (
	"github.com/insolar/insolar/instrumentation/insmetrics"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	depositCache = insmetrics.MustTagKey("deposit_cache")
)

var (
	depositCacheCount = stats.Int64(
		"deposit_cache_count",
		"count of deposit's cached utility records",
		stats.UnitDimensionless,
	)
)

func init() {
	err := view.Register(
		&view.View{
			Name:        depositCacheCount.Name(),
			Description: depositCacheCount.Description(),
			Measure:     depositCacheCount,
			Aggregation: view.LastValue(),
			TagKeys:     []tag.Key{depositCache},
		},
	)
	if err != nil {
		panic(err)
	}
}
