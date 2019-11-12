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

package api

import (
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-pg/pg/orm"

	"github.com/insolar/observer/internal/app/observer/postgres"

	"github.com/go-pg/pg"
	"github.com/insolar/insolar/insolar"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"

	"github.com/insolar/observer/component"
	"github.com/insolar/observer/internal/models"

	"github.com/sirupsen/logrus"
)

type Clock interface {
	Now() time.Time
}

type DefaultClock struct{}

func (c *DefaultClock) Now() time.Time {
	return time.Now()
}

type ObserverServer struct {
	db    *pg.DB
	log   *logrus.Logger
	clock Clock
	fee   *big.Int
}

func NewObserverServer(db *pg.DB, log *logrus.Logger, fee *big.Int, clock Clock) *ObserverServer {
	return &ObserverServer{db: db, log: log, clock: clock, fee: fee}
}

func (s *ObserverServer) GetMigrationAddresses(ctx echo.Context, params GetMigrationAddressesParams) error {
	limit := params.Limit
	if limit <= 0 || limit > 1000 {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError("`limit` should be in range [1, 1000]"))
	}

	query := s.db.Model(&models.MigrationAddress{}).
		Where("wasted = false")
	if params.Index != nil {
		id, err := strconv.ParseInt(*params.Index, 10, 64)
		if err != nil {
			s.log.Error(err)
			return ctx.JSON(http.StatusBadRequest, struct{}{})
		}
		query = query.Where("id > ?", id)
	}
	var result []models.MigrationAddress
	err := query.Order("id").Limit(limit).Select(&result)
	if err != nil {
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	resJSON := make([]interface{}, len(result))
	for i := 0; i < len(result); i++ {
		index := strconv.FormatInt(result[i].ID, 10)
		m := make(map[string]string, 2)
		m["address"] = result[i].Addr
		m["index"] = index
		resJSON[i] = m
	}
	return ctx.JSON(http.StatusOK, resJSON)
}

// GetMigrationAddressCount returns the total number of non-assigned migration addresses
func (s *ObserverServer) GetMigrationAddressCount(ctx echo.Context) error {
	count, err := s.db.Model(&models.MigrationAddress{}).
		Where("wasted = false").
		Count()
	if err != nil {
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	resJSON := make(map[string]int, 1)
	resJSON["count"] = count
	return ctx.JSON(http.StatusOK, resJSON)
}

func (s *ObserverServer) GetStatistics(ctx echo.Context) error {
	panic("implement me")
}

func (s *ObserverServer) TokenWeekPrice(ctx echo.Context, interval int) error {
	panic("implement me")
}

func (s *ObserverServer) TransactionsDetails(ctx echo.Context, txID string) error {
	panic("implement me")
}

// CloseTransactions returns a list of closed transactions (only with statuses `received` and `failed`).
func (s *ObserverServer) ClosedTransactions(ctx echo.Context, params ClosedTransactionsParams) error {
	limit := params.Limit
	if limit <= 0 || limit > 1000 {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError("`limit` should be in range [1, 1000]"))
	}

	var (
		pulseNumber    int64
		sequenceNumber int64
		err            error
	)
	if params.Index != nil {
		pulseNumber, sequenceNumber, err = checkIndex(*params.Index)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, NewSingleMessageError(err.Error()))
		}
	}

	var result []models.Transaction
	query := s.db.Model(&models.Transaction{}).
		Where("status_finished = ?", true)
	query, err = component.OrderByIndex(query, params.Order, pulseNumber, sequenceNumber, params.Index != nil, models.TxIndexTypeFinishPulseRecord)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError(err.Error()))
	}
	err = query.
		Limit(limit).
		Select(&result)
	if err != nil {
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	resJSON := make([]interface{}, len(result))
	for i := 0; i < len(result); i++ {
		resJSON[i] = TxToAPITx(result[i], models.TxIndexTypeFinishPulseRecord)
	}
	return ctx.JSON(http.StatusOK, resJSON)
}

func (s *ObserverServer) Fee(ctx echo.Context, amount string) error {
	return ctx.JSON(http.StatusOK, ResponsesFeeYaml{Fee: s.fee.String()})
}

func (s *ObserverServer) Member(ctx echo.Context, reference string) error {
	ref, errMsg := s.checkReference(reference)
	if errMsg != nil {
		return ctx.JSON(http.StatusBadRequest, *errMsg)
	}

	member, err := component.GetMember(ctx.Request().Context(), s.db, ref.Bytes())
	if err != nil {
		if err == component.ErrReferenceNotFound {
			return ctx.JSON(http.StatusNoContent, struct{}{})
		}
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	deposits, err := component.GetDeposits(ctx.Request().Context(), s.db, ref.Bytes())
	if err != nil {
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	return ctx.JSON(http.StatusOK, MemberToAPIMember(*member, *deposits, s.clock.Now().Unix()))
}

func (s *ObserverServer) Balance(ctx echo.Context, reference string) error {
	ref, errMsg := s.checkReference(reference)
	if errMsg != nil {
		return ctx.JSON(http.StatusBadRequest, *errMsg)
	}

	member, err := component.GetMemberBalance(ctx.Request().Context(), s.db, ref.Bytes())
	if err != nil {
		if err == component.ErrReferenceNotFound {
			return ctx.JSON(http.StatusNoContent, struct{}{})
		}
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	return ctx.JSON(http.StatusOK, ResponsesMemberBalanceYaml{Balance: member.Balance})
}

func (s *ObserverServer) MemberTransactions(ctx echo.Context, reference string, params MemberTransactionsParams) error {
	ref, errMsg := s.checkReference(reference)
	if errMsg != nil {
		return ctx.JSON(http.StatusBadRequest, *errMsg)
	}

	limit := params.Limit
	if limit <= 0 || limit > 1000 {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError("`limit` should be in range [1, 1000]"))
	}

	var errorMsg ErrorMessage

	var txs []models.Transaction
	query := s.db.Model(&txs)

	query, err := component.FilterByMemberReferenceAndDirection(query, ref, params.Direction)
	if err != nil {
		errorMsg.Error = append(errorMsg.Error, err.Error())
	}

	query = s.getTransactions(query, &errorMsg, params.Status, params.Type, params.Index, params.Order)

	if len(errorMsg.Error) > 0 {
		return ctx.JSON(http.StatusBadRequest, errorMsg)
	}

	err = query.Limit(limit).Select()

	if err != nil {
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	if len(txs) == 0 {
		return ctx.JSON(http.StatusNoContent, struct{}{})
	}

	res := SchemasTransactions{}
	for _, t := range txs {
		res = append(res, TxToAPITx(t, models.TxIndexTypePulseRecord))
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *ObserverServer) Notification(ctx echo.Context) error {
	panic("implement me")
}

func (s *ObserverServer) Transaction(ctx echo.Context, txIDStr string) error {
	txIDStr = strings.TrimSpace(txIDStr)

	if len(txIDStr) == 0 {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError("empty tx_id"))
	}

	txIDStr, err := url.QueryUnescape(txIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError("error unescaping tx_id parameter"))
	}

	txID, err := insolar.NewRecordReferenceFromString(txIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError("tx_id wrong format"))
	}

	tx, err := component.GetTx(ctx.Request().Context(), s.db, txID.Bytes())
	if err != nil {
		if err == component.ErrTxNotFound {
			return ctx.JSON(http.StatusNoContent, struct{}{})
		}
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	return ctx.JSON(http.StatusOK, TxToAPITx(*tx, models.TxIndexTypePulseRecord))
}

func (s *ObserverServer) TransactionsSearch(ctx echo.Context, params TransactionsSearchParams) error {
	limit := params.Limit
	if limit <= 0 || limit > 1000 {
		return ctx.JSON(http.StatusBadRequest, NewSingleMessageError("`limit` should be in range [1, 1000]"))
	}

	var errorMsg ErrorMessage
	var err error

	var txs []models.Transaction
	query := s.db.Model(&txs)

	if params.Value != nil {
		query, err = component.FilterByValue(query, *params.Value)
		if err != nil {
			errorMsg.Error = append(errorMsg.Error, err.Error())
		}
	}

	query = s.getTransactions(query, &errorMsg, params.Status, params.Type, params.Index, params.Order)
	if len(errorMsg.Error) > 0 {
		return ctx.JSON(http.StatusBadRequest, errorMsg)
	}
	err = query.Limit(params.Limit).Select()

	if err != nil {
		s.log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, struct{}{})
	}

	if len(txs) == 0 {
		return ctx.JSON(http.StatusNoContent, struct{}{})
	}

	res := SchemasTransactions{}
	for _, t := range txs {
		res = append(res, TxToAPITx(t, models.TxIndexTypePulseRecord))
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *ObserverServer) MarketStats(ctx echo.Context) error {
	panic("implement me")
}

func (s *ObserverServer) NetworkStats(ctx echo.Context) error {
	panic("implement me")
}

func (s *ObserverServer) SupplyStats(ctx echo.Context) error {
	repo := postgres.NewStatsRepository(s.db)
	xr := component.NewStatsManager(s.log, repo)
	result, err := xr.Coins()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "")
	}

	return ctx.JSON(http.StatusOK, ResponsesSupplyStatsYaml{
		TotalSupply:       result.Total(),
		MaxSupply:         result.Max(),
		CirculatingSupply: result.Circulating(),
	})
}

func (s *ObserverServer) SupplyStatsCirculating(ctx echo.Context) error {
	repo := postgres.NewStatsRepository(s.db)
	xr := component.NewStatsManager(s.log, repo)
	result, err := xr.Circulating()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "")
	}

	return ctx.String(http.StatusOK, result)
}

func (s *ObserverServer) SupplyStatsMax(ctx echo.Context) error {
	repo := postgres.NewStatsRepository(s.db)
	xr := component.NewStatsManager(s.log, repo)
	result, err := xr.Max()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "")
	}

	return ctx.String(http.StatusOK, result)
}

func (s *ObserverServer) SupplyStatsTotal(ctx echo.Context) error {
	repo := postgres.NewStatsRepository(s.db)
	xr := component.NewStatsManager(s.log, repo)
	result, err := xr.Total()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "")
	}

	return ctx.String(http.StatusOK, result)
}

func (s *ObserverServer) NewXNSTransferStats(ctx echo.Context) error {
	input := new(SchemaNewXNSMigration)
	if err := ctx.Bind(input); err != nil {
		return ctx.String(http.StatusBadRequest, "failed to unmarshal request body")
	}

	if len(input.Records) == 0 {
		return ctx.String(http.StatusBadRequest, "no records in input")
	}
	input.DaemonID = strings.TrimSpace(input.DaemonID)
	if len(input.DaemonID) == 0 {
		return ctx.String(http.StatusBadRequest, "daemonID isn't provided")
	}
	ref, err := s.checkReference(input.InsolarReference)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "wrong insolar reference")
	}

	migrationModels := []*postgres.MigrationStatsModel{}

	for _, rec := range input.Records {
		result, err := MigrationResultToModel(rec.Result)
		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}
		if len(rec.TxHash) == 0 {
			return ctx.String(http.StatusBadRequest, "empty tx hash")
		}

		xnsAmount := uint64(0)
		if rec.XnsAmount != nil {
			xnsAmount, err = strconv.ParseUint(*rec.XnsAmount, 10, 64)
			if err != nil {
				return ctx.String(
					http.StatusBadRequest,
					"failed to parse xns amount from"+*rec.XnsAmount,
				)
			}
		}

		if len(rec.BlockID) == 0 {
			return ctx.String(http.StatusBadRequest, "empty block id")
		}
		ethBlock, err := strconv.ParseUint(*rec.XnsAmount, 10, 64)
		if err != nil {
			return ctx.String(
				http.StatusBadRequest,
				"failed to parse block id from"+*rec.XnsAmount,
			)
		}

		model := &postgres.MigrationStatsModel{
			DaemonID:            input.DaemonID,
			InsolarRef:          ref.Bytes(),
			ModificationTime:    time.Time{},
			EthBlock:            ethBlock,
			TxHash:              rec.TxHash,
			Amount:              xnsAmount,
			Result:              result,
			ContractRequestBody: rec.ContractRequestBody,
			Error:               rec.Error,
		}

		migrationModels = append(migrationModels, model)
	}

	repo := postgres.NewMigrationStatsRepository(s.db)
	for _, model := range migrationModels {
		if err := repo.Insert(model); err != nil {
			s.log.Error(err)
			return ctx.String(http.StatusInternalServerError, "")
		}
	}

	return ctx.String(http.StatusOK, "")
}

func checkIndex(i string) (int64, int64, error) {
	index := strings.Split(i, ":")
	if len(index) != 2 {
		return 0, 0, errors.New("Query parameter 'index' should have the '<pulse_number>:<sequence_number>' format.") // nolint
	}
	var err error
	var pulseNumber, sequenceNumber int64
	pulseNumber, err = strconv.ParseInt(index[0], 10, 64)
	if err != nil {
		return 0, 0, errors.New("Query parameter 'index' should have the '<pulse_number>:<sequence_number>' format.") // nolint
	}
	sequenceNumber, err = strconv.ParseInt(index[1], 10, 64)
	if err != nil {
		return 0, 0, errors.New("Query parameter 'index' should have the '<pulse_number>:<sequence_number>' format.") // nolint
	}
	return pulseNumber, sequenceNumber, nil
}

func (s *ObserverServer) getTransactions(
	query *orm.Query, errorMsg *ErrorMessage, status, typeParam, index, order *string,
) *orm.Query {
	var err error
	if status != nil {
		query, err = component.FilterByStatus(query, *status)
		if err != nil {
			errorMsg.Error = append(errorMsg.Error, err.Error())
		}
	}

	if typeParam != nil {
		query, err = component.FilterByType(query, *typeParam)
		if err != nil {
			errorMsg.Error = append(errorMsg.Error, err.Error())
		}
	}

	var pulseNumber int64
	var sequenceNumber int64
	byIndex := false
	if index != nil {
		pulseNumber, sequenceNumber, err = checkIndex(*index)
		if err != nil {
			errorMsg.Error = append(errorMsg.Error, err.Error())
		} else {
			byIndex = true
		}
	}

	query, err = component.OrderByIndex(query, order, pulseNumber, sequenceNumber, byIndex, models.TxIndexTypePulseRecord)
	if err != nil {
		errorMsg.Error = append(errorMsg.Error, err.Error())
	}
	return query
}

func (s *ObserverServer) checkReference(referenceRow string) (*insolar.Reference, *ErrorMessage) {
	referenceRow = strings.TrimSpace(referenceRow)
	var errMsg ErrorMessage

	if len(referenceRow) == 0 {
		errMsg = NewSingleMessageError("empty reference")
		return nil, &errMsg
	}

	reference, err := url.QueryUnescape(referenceRow)
	if err != nil {
		errMsg = NewSingleMessageError("error unescaping reference parameter")
		return nil, &errMsg
	}

	ref, err := insolar.NewReferenceFromString(reference)
	if err != nil {
		errMsg = NewSingleMessageError("reference wrong format")
		return nil, &errMsg
	}
	return ref, nil
}
