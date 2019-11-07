package collecting

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/insolar/insolar/application/appfoundation"
	"github.com/insolar/insolar/application/builtin/contract/member"
	"github.com/insolar/insolar/application/builtin/contract/member/signer"
	proxyDeposit "github.com/insolar/insolar/application/builtin/proxy/deposit"
	proxyMember "github.com/insolar/insolar/application/builtin/proxy/member"
	"github.com/insolar/insolar/application/genesisrefs"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/ledger/heavy/exporter"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/insolar/observer/internal/app/observer"
	"github.com/insolar/observer/internal/app/observer/store"
	"github.com/insolar/observer/internal/models"
)

const (
	callSiteTransfer  = "member.transfer"
	callSiteMigration = "deposit.migration"
	callSiteRelease   = "deposit.transfer"
)

const (
	methodCall              = "Call"
	methodTransferToDeposit = "TransferToDeposit"
	methodTransfer          = "Transfer"
	methodAccept            = "Accept"
)

const (
	paramAmount      = "amount"
	paramToMemberRef = "toMemberReference"
)

type TxRegisterCollector struct {
	log *logrus.Logger
}

func NewTxRegisterCollector(log *logrus.Logger) *TxRegisterCollector {
	return &TxRegisterCollector{
		log: log,
	}
}

func (c *TxRegisterCollector) Collect(ctx context.Context, rec exporter.Record) *observer.TxRegister {
	log := c.log.WithField("record_id", rec.Record.ID.DebugString())
	log = log.WithField("collect_process_id", uuid.New())
	log.Debug("received record")
	defer log.Debug("record processed")

	request, ok := record.Unwrap(&rec.Record.Virtual).(*record.IncomingRequest)
	if !ok {
		log.Debug("skipped (not IncomingRequest)")
		return nil
	}

	log.Debug("parsing method ", request.Method)
	var tx *observer.TxRegister
	switch request.Method {
	case methodCall:
		tx = c.fromTransfer(log, rec)
	case methodTransferToDeposit:
		tx = c.fromMigration(log, rec)
	case methodTransfer:
		tx = c.fromRelease(log, rec)
	default:
		return nil
	}
	if tx == nil {
		return nil
	}
	if err := tx.Validate(); err != nil {
		log.Error(errors.Wrap(err, "invalid transaction received"))
		return nil
	}
	return tx
}

func (c *TxRegisterCollector) fromTransfer(log *logrus.Entry, rec exporter.Record) *observer.TxRegister {
	txID := *insolar.NewRecordReference(rec.Record.ID)
	log = log.WithField("tx_id", txID.GetLocal().DebugString())

	request, ok := record.Unwrap(&rec.Record.Virtual).(*record.IncomingRequest)
	if !ok {
		log.Debug("skipped (not IncomingRequest)")
		return nil
	}

	// Skip non-member objects.
	if request.Prototype != nil && !request.Prototype.Equal(*proxyMember.PrototypeReference) {
		log.Debugf("skipped (not member object)")
		return nil
	}

	if request.Method != methodCall {
		log.Debug("skipped (not Call method)")
		return nil
	}

	// Skip internal calls.
	if request.APINode.IsEmpty() {
		log.Debug("skipped (APINode is empty)")
		return nil
	}

	// Skip saga.
	if request.IsDetachedCall() {
		log.Debug("skipped (saga)")
		return nil
	}

	args, callParams, err := parseExternalArguments(request.Arguments)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse arguments"))
		return nil
	}
	if args.Params.CallSite != callSiteTransfer {
		log.Debug("skipped (CallSite not callSiteTransfer)")
		return nil
	}

	memberFrom, err := insolar.NewObjectReferenceFromString(args.Params.Reference)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse from reference"))
		return nil
	}
	toMemberStr, ok := callParams[paramToMemberRef].(string)
	if !ok {
		log.Error(errors.Wrap(err, "failed to parse from reference"))
		return nil
	}
	memberTo, err := insolar.NewObjectReferenceFromString(toMemberStr)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse to reference"))
		return nil
	}
	amount, ok := callParams[paramAmount].(string)
	if !ok {
		log.Error(errors.Wrap(err, "failed to parse from amount"))
		return nil
	}

	log.Debug("created TxRegister")
	return &observer.TxRegister{
		Type:                models.TTypeTransfer,
		TransactionID:       txID,
		PulseNumber:         int64(rec.Record.ID.Pulse()),
		RecordNumber:        int64(rec.RecordNumber),
		Amount:              amount,
		MemberFromReference: memberFrom.Bytes(),
		MemberToReference:   memberTo.Bytes(),
	}
}

func (c *TxRegisterCollector) fromMigration(log *logrus.Entry, rec exporter.Record) *observer.TxRegister {
	request, ok := record.Unwrap(&rec.Record.Virtual).(*record.IncomingRequest)
	if !ok {
		log.Debug("skipped (not IncomingRequest)")
		return nil
	}

	// Skip non-deposit objects.
	if request.Prototype == nil || !request.Prototype.Equal(*proxyDeposit.PrototypeReference) {
		log.Debug("skipped (not deposit object)")
		return nil
	}

	if request.Method != methodTransferToDeposit {
		log.Debug("skipped (not TransferToDeposit method)")
		return nil
	}

	// Skip external calls.
	if request.Caller.IsEmpty() {
		log.Debug("skipped (Caller is empty)")
		return nil
	}

	var (
		amount                                string
		txID, toDeposit, fromMember, toMember insolar.Reference
	)
	err := insolar.Deserialize(request.Arguments, []interface{}{&amount, &toDeposit, &fromMember, &txID, &toMember})
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse arguments"))
		return nil
	}

	// Ensure txID is record reference so other collectors can match it.
	txID = *insolar.NewRecordReference(*txID.GetLocal())

	log = log.WithField("tx_id", txID.GetLocal().DebugString())
	log.Debug("created TxRegister")
	return &observer.TxRegister{
		Type:                models.TTypeMigration,
		TransactionID:       txID,
		PulseNumber:         int64(rec.Record.ID.Pulse()),
		RecordNumber:        int64(rec.RecordNumber),
		MemberFromReference: fromMember.Bytes(),
		MemberToReference:   toMember.Bytes(),
		DepositToReference:  toDeposit.Bytes(),
		Amount:              amount,
	}
}

func (c *TxRegisterCollector) fromRelease(log *logrus.Entry, rec exporter.Record) *observer.TxRegister {
	request, ok := record.Unwrap(&rec.Record.Virtual).(*record.IncomingRequest)
	if !ok {
		log.Debug("skipped (not IncomingRequest)")
		return nil
	}

	// Skip non-deposit objects.
	if request.Prototype == nil || !request.Prototype.Equal(*proxyDeposit.PrototypeReference) {
		log.Debug("skipped (not deposit object)")
		return nil
	}

	if request.Method != methodTransfer {
		log.Debug("skipped (not Transfer method)")
		return nil
	}

	// Skip external calls.
	if request.Caller.IsEmpty() {
		log.Debug("skipped (Caller is empty)")
		return nil
	}

	var (
		amount         string
		txID, toMember insolar.Reference
	)
	err := insolar.Deserialize(request.Arguments, []interface{}{&amount, &toMember, &txID})
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse arguments"))
		return nil
	}

	// Ensure txID is record reference so other collectors can match it.
	txID = *insolar.NewRecordReference(*txID.GetLocal())

	log = log.WithField("tx_id", txID.GetLocal().DebugString())
	log.Debug("created TxRegister")
	return &observer.TxRegister{
		Type:                 models.TTypeRelease,
		TransactionID:        txID,
		PulseNumber:          int64(rec.Record.ID.Pulse()),
		RecordNumber:         int64(rec.RecordNumber),
		MemberToReference:    toMember.Bytes(),
		DepositFromReference: insolar.NewReference(rec.Record.ObjectID).Bytes(),
		Amount:               amount,
	}
}

func parseExternalArguments(in []byte) (member.Request, map[string]interface{}, error) {
	if in == nil {
		return member.Request{}, nil, nil
	}
	var signedRequest []byte
	err := insolar.Deserialize(in, []interface{}{&signedRequest})
	if err != nil {
		return member.Request{}, nil, err
	}

	if len(signedRequest) == 0 {
		return member.Request{}, nil, errors.New("failed to parse signed request")
	}
	request := member.Request{}
	{
		var encodedRequest []byte
		// IMPORTANT: argument number should match serialization. This is why we use nil as second and third values.
		err = signer.UnmarshalParams(signedRequest, []interface{}{&encodedRequest, nil, nil}...)
		if err != nil {
			return member.Request{}, nil, errors.Wrapf(err, "failed to unmarshal params")
		}
		err = json.Unmarshal(encodedRequest, &request)
		if err != nil {
			return member.Request{}, nil, errors.Wrapf(err, "failed to unmarshal json member request")
		}
	}

	if request.Params.CallParams == nil {
		return request, nil, nil
	}

	callParams, ok := request.Params.CallParams.(map[string]interface{})
	if !ok {
		return member.Request{}, nil, errors.New("failed to decode CallParams")
	}
	return request, callParams, nil
}

type TxResultCollector struct {
	fetcher store.RecordFetcher
	log     *logrus.Logger
}

func NewTxResultCollector(log *logrus.Logger, fetcher store.RecordFetcher) *TxResultCollector {
	return &TxResultCollector{
		fetcher: fetcher,
		log:     log,
	}
}

func (c *TxResultCollector) Collect(ctx context.Context, rec exporter.Record) *observer.TxResult {
	log := c.log.WithField("record_id", rec.Record.ID.DebugString())
	log = log.WithField("collect_process_id", uuid.New())
	log.Debug("received record")
	defer log.Debug("record processed")

	result, ok := record.Unwrap(&rec.Record.Virtual).(*record.Result)
	if !ok {
		log.Debug("skipped (not Result)")
		return nil
	}

	// Ensure txID is record reference so other collectors can match it.
	txID := *insolar.NewRecordReference(*result.Request.GetLocal())
	log = log.WithField("tx_id", txID.GetLocal().DebugString())

	requestRecord, err := c.fetcher.Request(ctx, *txID.GetLocal())
	if err != nil {
		log.Error(errors.Wrapf(
			err,
			"failed to fetch request with id %s",
			txID.GetLocal().DebugString()),
		)
		return nil
	}

	request, ok := record.Unwrap(&requestRecord.Virtual).(*record.IncomingRequest)
	if !ok {
		log.Debug("skipped (matching request is not IncomingRequest)")
		return nil
	}

	if request.Method != methodCall {
		log.Debug("skipped (method is not Call)")
		return nil
	}
	// Skip non-API requests.
	if request.APINode.IsEmpty() {
		log.Debug("skipped (APINode is empty)")
		return nil
	}
	// Skip saga.
	if request.IsDetachedCall() {
		log.Debug("skipped (request is saga)")
		return nil
	}
	args, _, err := parseExternalArguments(request.Arguments)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse request arguments"))
		return nil
	}

	// Migration and release don't have fees.
	if args.Params.CallSite == callSiteMigration || args.Params.CallSite == callSiteRelease {
		tx := &observer.TxResult{
			TransactionID: txID,
			Fee:           "0",
		}
		if err = tx.Validate(); err != nil {
			log.Error(errors.Wrap(err, "failed to validate transaction"))
			return nil
		}
		return tx
	}

	// Processing transfer between members. Its the only transfer that has fee.
	if args.Params.CallSite != callSiteTransfer {
		log.Debug("skipped (callSite is not Transfer)")
		return nil
	}
	response := member.TransferResponse{}
	err = insolar.Deserialize(result.Payload, &foundation.Result{
		Returns: []interface{}{&response, nil},
	})
	if err != nil {
		log.Error(errors.Wrap(err, "failed to deserialize method result"))
		return nil
	}

	log = log.WithField("tx_id", txID.GetLocal().DebugString())
	log.Debug("created TxResult")
	tx := &observer.TxResult{
		TransactionID: txID,
		Fee:           response.Fee,
	}
	if err = tx.Validate(); err != nil {
		log.Error(errors.Wrap(err, "failed to validate transaction"))
		return nil
	}
	return tx
}

type TxSagaResultCollector struct {
	fetcher store.RecordFetcher
	log     *logrus.Logger
}

func NewTxSagaResultCollector(log *logrus.Logger, fetcher store.RecordFetcher) *TxSagaResultCollector {
	return &TxSagaResultCollector{
		fetcher: fetcher,
		log:     log,
	}
}

func (c *TxSagaResultCollector) Collect(ctx context.Context, rec exporter.Record) *observer.TxSagaResult {
	log := c.log.WithField("record_id", rec.Record.ID.DebugString())
	log = log.WithField("collect_process_id", uuid.New())
	log.Debug("received record")
	defer log.Debug("record processed")

	if rec.Record.ObjectID == *genesisrefs.ContractFeeMember.GetLocal() {
		log.Debug("skipped (fee member object)")
		return nil
	}

	result, ok := record.Unwrap(&rec.Record.Virtual).(*record.Result)
	if !ok {
		log.Debug("skipped (not Result)")
		return nil
	}

	log = log.WithField("request_id", result.Request.GetLocal().DebugString())

	requestRecord, err := c.fetcher.Request(ctx, *result.Request.GetLocal())
	if err != nil {
		log.Error(errors.Wrapf(
			err,
			"failed to fetch request with id %s",
			result.Request.GetLocal().DebugString()),
		)
		return nil
	}

	request, ok := record.Unwrap(&requestRecord.Virtual).(*record.IncomingRequest)
	if !ok {
		return nil
	}

	log.Debug("parsing method ", request.Method)
	var tx *observer.TxSagaResult
	switch request.Method {
	case methodAccept:
		tx = c.fromAccept(log, rec, *request, *result)
	case methodCall:
		tx = c.fromCall(log, rec, *request, *result)
	}
	if tx != nil {
		if err := tx.Validate(); err != nil {
			log.Error(errors.Wrap(err, "failed to validate transaction"))
			return nil
		}
	}
	return tx
}

func (c *TxSagaResultCollector) fromAccept(
	log *logrus.Entry,
	resultRec exporter.Record,
	request record.IncomingRequest,
	result record.Result,
) *observer.TxSagaResult {
	// Skip non-saga.
	if !request.IsDetachedCall() {
		log.Debug("skipped (request is not saga)")
		return nil
	}

	var acceptArgs appfoundation.SagaAcceptInfo
	err := insolar.Deserialize(request.Arguments, []interface{}{&acceptArgs})
	if err != nil {
		log.Error(errors.Wrap(err, "failed to deserialize method arguments"))
		return nil
	}
	// Ensure txID is record reference so other collectors can match it.
	txID := *insolar.NewRecordReference(*acceptArgs.Request.GetLocal())
	log = log.WithField("tx_id", txID.GetLocal().DebugString())

	response := foundation.Result{}
	err = insolar.Deserialize(result.Payload, &response)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to deserialize method result"))
		return nil
	}

	if len(response.Returns) < 1 {
		log.Error(errors.Wrap(err, "unexpected number of Accept method returned parameters"))
		return nil
	}

	// The first return parameter of Accept method is error, so we check if its not nil.
	if response.Error != nil || response.Returns[0] != nil {
		log.Error("saga resulted with error")
		log.Debug("created failed TxSagaResult")
		return &observer.TxSagaResult{
			TransactionID:      txID,
			FinishSuccess:      false,
			FinishPulseNumber:  int64(resultRec.Record.ID.Pulse()),
			FinishRecordNumber: int64(resultRec.RecordNumber),
		}
	}

	log.Debug("created success TxSagaResult")
	return &observer.TxSagaResult{
		TransactionID:      txID,
		FinishSuccess:      true,
		FinishPulseNumber:  int64(resultRec.Record.ID.Pulse()),
		FinishRecordNumber: int64(resultRec.RecordNumber),
	}
}

func (c *TxSagaResultCollector) fromCall(
	log *logrus.Entry,
	resultRec exporter.Record,
	request record.IncomingRequest,
	result record.Result,
) *observer.TxSagaResult {

	// Ensure txID is record reference so other collectors can match it.
	txID := *insolar.NewRecordReference(*result.Request.GetLocal())
	log = log.WithField("tx_id", txID.GetLocal().DebugString())

	// Skip non-API requests.
	if request.APINode.IsEmpty() {
		log.Debug("skipped (request APINode is empty)")
		return nil
	}
	// Skip saga.
	if request.IsDetachedCall() {
		log.Debug("skipped (request is saga)")
		return nil
	}
	args, _, err := parseExternalArguments(request.Arguments)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to parse request arguments"))
		return nil
	}

	isTransfer := args.Params.CallSite == callSiteTransfer
	isMigration := args.Params.CallSite == callSiteMigration
	isRelease := args.Params.CallSite == callSiteRelease
	if !isTransfer && !isMigration && !isRelease {
		log.Debug("skipped (request callSite is not parsable)")
		return nil
	}

	var response foundation.Result
	err = insolar.Deserialize(result.Payload, &response)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to deserialize method result"))
		return nil
	}
	if len(response.Returns) < 2 {
		log.Error(errors.Wrap(err, "unexpected number of Call method returned parameters"))
		return nil
	}

	// The second return parameter of Call method is error, so we check if its not nil.
	if response.Error != nil || response.Returns[1] != nil {
		log.Debug("created failed TxSagaResult")
		return &observer.TxSagaResult{
			TransactionID:      txID,
			FinishSuccess:      false,
			FinishPulseNumber:  int64(resultRec.Record.ID.Pulse()),
			FinishRecordNumber: int64(resultRec.RecordNumber),
		}
	}

	// Successful call does not produce transactions since it will be produced by saga call. It avoids double insert
	// on conflict.
	return nil
}
