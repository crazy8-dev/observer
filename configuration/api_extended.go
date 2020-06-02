// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer/blob/master/LICENSE.md.

// +build !public

package configuration

import (
	"math/big"
	"time"
)

type APIExtended struct {
	API `mapstructure:",squash"`

	FeeAmount            *big.Int
	PriceOrigin          string
	Price                string
	CMCMarketStatsParams CMCMarketStatsParamsEnabled
}

func (APIExtended) Default() *APIExtended {
	return &APIExtended{
		API: API{
			Listen: ":0",
			DB: DB{
				URL:             "postgres://postgres@localhost/postgres?sslmode=disable",
				Attempts:        5,
				AttemptInterval: 3 * time.Second,
			},
			Log: Log{
				Level:        "debug",
				Format:       "text",
				OutputType:   "stderr",
				OutputParams: "",
				Buffer:       0,
			},
		},
		FeeAmount:   big.NewInt(1000000000),
		Price:       "0.05",
		PriceOrigin: "const", // const|binance|coin_market_cap
		CMCMarketStatsParams: CMCMarketStatsParamsEnabled{
			CirculatingSupply: true,
			DailyChange:       true,
			MarketCap:         true,
			Rank:              true,
			Volume:            true,
		},
	}
}

func (a APIExtended) GetListen() string {
	return a.Listen
}

func (a APIExtended) GetDB() DB {
	return a.DB
}

func (a APIExtended) GetLog() Log {
	return a.Log
}

func (a APIExtended) GetFeeAmount() *big.Int {
	return a.FeeAmount
}

func (a APIExtended) GetPriceOrigin() string {
	return a.PriceOrigin
}

func (a APIExtended) GetPrice() string {
	return a.Price
}

func (a APIExtended) GetCMCMarketStatsParams() CMCMarketStatsParamsEnabled {
	return a.CMCMarketStatsParams
}

func GetAPIConfig() APIConfig {
	return &APIExtended{}
}
