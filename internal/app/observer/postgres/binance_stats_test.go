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

package postgres_test

import (
	"testing"

	"github.com/insolar/observer/internal/app/observer/postgres"
	"github.com/insolar/observer/internal/models"
	"github.com/stretchr/testify/require"
)

func TestBinanceStats(t *testing.T) {
	repo := postgres.NewBinanceStatsRepository(db)

	_, err := repo.LastStats()
	require.Error(t, err)
	require.Contains(t, err.Error(), "no data")

	stat := models.BinanceStats{
		Symbol:             "TEST",
		SymbolPriceBTC:     "111",
		SymbolPriceUSD:     "2222",
		BTCPriceUSD:        "3333",
		PriceChangePercent: "4444",
	}
	err = repo.InsertStats(&stat)
	require.NoError(t, err)

	savedData, err := repo.LastStats()
	require.NoError(t, err)

	require.Equal(t, stat.PriceChangePercent, savedData.PriceChangePercent)
	require.Equal(t, stat.BTCPriceUSD, savedData.BTCPriceUSD)
	require.Equal(t, stat.SymbolPriceUSD, savedData.SymbolPriceUSD)
	require.Equal(t, stat.SymbolPriceBTC, savedData.SymbolPriceBTC)
	require.Equal(t, stat.Symbol, savedData.Symbol)
}
