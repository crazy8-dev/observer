package object

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
)

// RecordStorageMock implements RecordStorage
type RecordStorageMock struct {
	t minimock.Tester

	funcBatchSet          func(ctx context.Context, recs []record.Material) (err error)
	inspectFuncBatchSet   func(ctx context.Context, recs []record.Material)
	afterBatchSetCounter  uint64
	beforeBatchSetCounter uint64
	BatchSetMock          mRecordStorageMockBatchSet

	funcForID          func(ctx context.Context, id insolar.ID) (m1 record.Material, err error)
	inspectFuncForID   func(ctx context.Context, id insolar.ID)
	afterForIDCounter  uint64
	beforeForIDCounter uint64
	ForIDMock          mRecordStorageMockForID

	funcSet          func(ctx context.Context, rec record.Material) (err error)
	inspectFuncSet   func(ctx context.Context, rec record.Material)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mRecordStorageMockSet
}

// NewRecordStorageMock returns a mock for RecordStorage
func NewRecordStorageMock(t minimock.Tester) *RecordStorageMock {
	m := &RecordStorageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.BatchSetMock = mRecordStorageMockBatchSet{mock: m}
	m.BatchSetMock.callArgs = []*RecordStorageMockBatchSetParams{}

	m.ForIDMock = mRecordStorageMockForID{mock: m}
	m.ForIDMock.callArgs = []*RecordStorageMockForIDParams{}

	m.SetMock = mRecordStorageMockSet{mock: m}
	m.SetMock.callArgs = []*RecordStorageMockSetParams{}

	return m
}

type mRecordStorageMockBatchSet struct {
	mock               *RecordStorageMock
	defaultExpectation *RecordStorageMockBatchSetExpectation
	expectations       []*RecordStorageMockBatchSetExpectation

	callArgs []*RecordStorageMockBatchSetParams
	mutex    sync.RWMutex
}

// RecordStorageMockBatchSetExpectation specifies expectation struct of the RecordStorage.BatchSet
type RecordStorageMockBatchSetExpectation struct {
	mock    *RecordStorageMock
	params  *RecordStorageMockBatchSetParams
	results *RecordStorageMockBatchSetResults
	Counter uint64
}

// RecordStorageMockBatchSetParams contains parameters of the RecordStorage.BatchSet
type RecordStorageMockBatchSetParams struct {
	ctx  context.Context
	recs []record.Material
}

// RecordStorageMockBatchSetResults contains results of the RecordStorage.BatchSet
type RecordStorageMockBatchSetResults struct {
	err error
}

// Expect sets up expected params for RecordStorage.BatchSet
func (mmBatchSet *mRecordStorageMockBatchSet) Expect(ctx context.Context, recs []record.Material) *mRecordStorageMockBatchSet {
	if mmBatchSet.mock.funcBatchSet != nil {
		mmBatchSet.mock.t.Fatalf("RecordStorageMock.BatchSet mock is already set by Set")
	}

	if mmBatchSet.defaultExpectation == nil {
		mmBatchSet.defaultExpectation = &RecordStorageMockBatchSetExpectation{}
	}

	mmBatchSet.defaultExpectation.params = &RecordStorageMockBatchSetParams{ctx, recs}
	for _, e := range mmBatchSet.expectations {
		if minimock.Equal(e.params, mmBatchSet.defaultExpectation.params) {
			mmBatchSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmBatchSet.defaultExpectation.params)
		}
	}

	return mmBatchSet
}

// Inspect accepts an inspector function that has same arguments as the RecordStorage.BatchSet
func (mmBatchSet *mRecordStorageMockBatchSet) Inspect(f func(ctx context.Context, recs []record.Material)) *mRecordStorageMockBatchSet {
	if mmBatchSet.mock.inspectFuncBatchSet != nil {
		mmBatchSet.mock.t.Fatalf("Inspect function is already set for RecordStorageMock.BatchSet")
	}

	mmBatchSet.mock.inspectFuncBatchSet = f

	return mmBatchSet
}

// Return sets up results that will be returned by RecordStorage.BatchSet
func (mmBatchSet *mRecordStorageMockBatchSet) Return(err error) *RecordStorageMock {
	if mmBatchSet.mock.funcBatchSet != nil {
		mmBatchSet.mock.t.Fatalf("RecordStorageMock.BatchSet mock is already set by Set")
	}

	if mmBatchSet.defaultExpectation == nil {
		mmBatchSet.defaultExpectation = &RecordStorageMockBatchSetExpectation{mock: mmBatchSet.mock}
	}
	mmBatchSet.defaultExpectation.results = &RecordStorageMockBatchSetResults{err}
	return mmBatchSet.mock
}

//Set uses given function f to mock the RecordStorage.BatchSet method
func (mmBatchSet *mRecordStorageMockBatchSet) Set(f func(ctx context.Context, recs []record.Material) (err error)) *RecordStorageMock {
	if mmBatchSet.defaultExpectation != nil {
		mmBatchSet.mock.t.Fatalf("Default expectation is already set for the RecordStorage.BatchSet method")
	}

	if len(mmBatchSet.expectations) > 0 {
		mmBatchSet.mock.t.Fatalf("Some expectations are already set for the RecordStorage.BatchSet method")
	}

	mmBatchSet.mock.funcBatchSet = f
	return mmBatchSet.mock
}

// When sets expectation for the RecordStorage.BatchSet which will trigger the result defined by the following
// Then helper
func (mmBatchSet *mRecordStorageMockBatchSet) When(ctx context.Context, recs []record.Material) *RecordStorageMockBatchSetExpectation {
	if mmBatchSet.mock.funcBatchSet != nil {
		mmBatchSet.mock.t.Fatalf("RecordStorageMock.BatchSet mock is already set by Set")
	}

	expectation := &RecordStorageMockBatchSetExpectation{
		mock:   mmBatchSet.mock,
		params: &RecordStorageMockBatchSetParams{ctx, recs},
	}
	mmBatchSet.expectations = append(mmBatchSet.expectations, expectation)
	return expectation
}

// Then sets up RecordStorage.BatchSet return parameters for the expectation previously defined by the When method
func (e *RecordStorageMockBatchSetExpectation) Then(err error) *RecordStorageMock {
	e.results = &RecordStorageMockBatchSetResults{err}
	return e.mock
}

// BatchSet implements RecordStorage
func (mmBatchSet *RecordStorageMock) BatchSet(ctx context.Context, recs []record.Material) (err error) {
	mm_atomic.AddUint64(&mmBatchSet.beforeBatchSetCounter, 1)
	defer mm_atomic.AddUint64(&mmBatchSet.afterBatchSetCounter, 1)

	if mmBatchSet.inspectFuncBatchSet != nil {
		mmBatchSet.inspectFuncBatchSet(ctx, recs)
	}

	mm_params := &RecordStorageMockBatchSetParams{ctx, recs}

	// Record call args
	mmBatchSet.BatchSetMock.mutex.Lock()
	mmBatchSet.BatchSetMock.callArgs = append(mmBatchSet.BatchSetMock.callArgs, mm_params)
	mmBatchSet.BatchSetMock.mutex.Unlock()

	for _, e := range mmBatchSet.BatchSetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmBatchSet.BatchSetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmBatchSet.BatchSetMock.defaultExpectation.Counter, 1)
		mm_want := mmBatchSet.BatchSetMock.defaultExpectation.params
		mm_got := RecordStorageMockBatchSetParams{ctx, recs}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmBatchSet.t.Errorf("RecordStorageMock.BatchSet got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmBatchSet.BatchSetMock.defaultExpectation.results
		if mm_results == nil {
			mmBatchSet.t.Fatal("No results are set for the RecordStorageMock.BatchSet")
		}
		return (*mm_results).err
	}
	if mmBatchSet.funcBatchSet != nil {
		return mmBatchSet.funcBatchSet(ctx, recs)
	}
	mmBatchSet.t.Fatalf("Unexpected call to RecordStorageMock.BatchSet. %v %v", ctx, recs)
	return
}

// BatchSetAfterCounter returns a count of finished RecordStorageMock.BatchSet invocations
func (mmBatchSet *RecordStorageMock) BatchSetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBatchSet.afterBatchSetCounter)
}

// BatchSetBeforeCounter returns a count of RecordStorageMock.BatchSet invocations
func (mmBatchSet *RecordStorageMock) BatchSetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBatchSet.beforeBatchSetCounter)
}

// Calls returns a list of arguments used in each call to RecordStorageMock.BatchSet.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmBatchSet *mRecordStorageMockBatchSet) Calls() []*RecordStorageMockBatchSetParams {
	mmBatchSet.mutex.RLock()

	argCopy := make([]*RecordStorageMockBatchSetParams, len(mmBatchSet.callArgs))
	copy(argCopy, mmBatchSet.callArgs)

	mmBatchSet.mutex.RUnlock()

	return argCopy
}

// MinimockBatchSetDone returns true if the count of the BatchSet invocations corresponds
// the number of defined expectations
func (m *RecordStorageMock) MinimockBatchSetDone() bool {
	for _, e := range m.BatchSetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.BatchSetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterBatchSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBatchSet != nil && mm_atomic.LoadUint64(&m.afterBatchSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockBatchSetInspect logs each unmet expectation
func (m *RecordStorageMock) MinimockBatchSetInspect() {
	for _, e := range m.BatchSetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RecordStorageMock.BatchSet with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.BatchSetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterBatchSetCounter) < 1 {
		if m.BatchSetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RecordStorageMock.BatchSet")
		} else {
			m.t.Errorf("Expected call to RecordStorageMock.BatchSet with params: %#v", *m.BatchSetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBatchSet != nil && mm_atomic.LoadUint64(&m.afterBatchSetCounter) < 1 {
		m.t.Error("Expected call to RecordStorageMock.BatchSet")
	}
}

type mRecordStorageMockForID struct {
	mock               *RecordStorageMock
	defaultExpectation *RecordStorageMockForIDExpectation
	expectations       []*RecordStorageMockForIDExpectation

	callArgs []*RecordStorageMockForIDParams
	mutex    sync.RWMutex
}

// RecordStorageMockForIDExpectation specifies expectation struct of the RecordStorage.ForID
type RecordStorageMockForIDExpectation struct {
	mock    *RecordStorageMock
	params  *RecordStorageMockForIDParams
	results *RecordStorageMockForIDResults
	Counter uint64
}

// RecordStorageMockForIDParams contains parameters of the RecordStorage.ForID
type RecordStorageMockForIDParams struct {
	ctx context.Context
	id  insolar.ID
}

// RecordStorageMockForIDResults contains results of the RecordStorage.ForID
type RecordStorageMockForIDResults struct {
	m1  record.Material
	err error
}

// Expect sets up expected params for RecordStorage.ForID
func (mmForID *mRecordStorageMockForID) Expect(ctx context.Context, id insolar.ID) *mRecordStorageMockForID {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("RecordStorageMock.ForID mock is already set by Set")
	}

	if mmForID.defaultExpectation == nil {
		mmForID.defaultExpectation = &RecordStorageMockForIDExpectation{}
	}

	mmForID.defaultExpectation.params = &RecordStorageMockForIDParams{ctx, id}
	for _, e := range mmForID.expectations {
		if minimock.Equal(e.params, mmForID.defaultExpectation.params) {
			mmForID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForID.defaultExpectation.params)
		}
	}

	return mmForID
}

// Inspect accepts an inspector function that has same arguments as the RecordStorage.ForID
func (mmForID *mRecordStorageMockForID) Inspect(f func(ctx context.Context, id insolar.ID)) *mRecordStorageMockForID {
	if mmForID.mock.inspectFuncForID != nil {
		mmForID.mock.t.Fatalf("Inspect function is already set for RecordStorageMock.ForID")
	}

	mmForID.mock.inspectFuncForID = f

	return mmForID
}

// Return sets up results that will be returned by RecordStorage.ForID
func (mmForID *mRecordStorageMockForID) Return(m1 record.Material, err error) *RecordStorageMock {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("RecordStorageMock.ForID mock is already set by Set")
	}

	if mmForID.defaultExpectation == nil {
		mmForID.defaultExpectation = &RecordStorageMockForIDExpectation{mock: mmForID.mock}
	}
	mmForID.defaultExpectation.results = &RecordStorageMockForIDResults{m1, err}
	return mmForID.mock
}

//Set uses given function f to mock the RecordStorage.ForID method
func (mmForID *mRecordStorageMockForID) Set(f func(ctx context.Context, id insolar.ID) (m1 record.Material, err error)) *RecordStorageMock {
	if mmForID.defaultExpectation != nil {
		mmForID.mock.t.Fatalf("Default expectation is already set for the RecordStorage.ForID method")
	}

	if len(mmForID.expectations) > 0 {
		mmForID.mock.t.Fatalf("Some expectations are already set for the RecordStorage.ForID method")
	}

	mmForID.mock.funcForID = f
	return mmForID.mock
}

// When sets expectation for the RecordStorage.ForID which will trigger the result defined by the following
// Then helper
func (mmForID *mRecordStorageMockForID) When(ctx context.Context, id insolar.ID) *RecordStorageMockForIDExpectation {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("RecordStorageMock.ForID mock is already set by Set")
	}

	expectation := &RecordStorageMockForIDExpectation{
		mock:   mmForID.mock,
		params: &RecordStorageMockForIDParams{ctx, id},
	}
	mmForID.expectations = append(mmForID.expectations, expectation)
	return expectation
}

// Then sets up RecordStorage.ForID return parameters for the expectation previously defined by the When method
func (e *RecordStorageMockForIDExpectation) Then(m1 record.Material, err error) *RecordStorageMock {
	e.results = &RecordStorageMockForIDResults{m1, err}
	return e.mock
}

// ForID implements RecordStorage
func (mmForID *RecordStorageMock) ForID(ctx context.Context, id insolar.ID) (m1 record.Material, err error) {
	mm_atomic.AddUint64(&mmForID.beforeForIDCounter, 1)
	defer mm_atomic.AddUint64(&mmForID.afterForIDCounter, 1)

	if mmForID.inspectFuncForID != nil {
		mmForID.inspectFuncForID(ctx, id)
	}

	mm_params := &RecordStorageMockForIDParams{ctx, id}

	// Record call args
	mmForID.ForIDMock.mutex.Lock()
	mmForID.ForIDMock.callArgs = append(mmForID.ForIDMock.callArgs, mm_params)
	mmForID.ForIDMock.mutex.Unlock()

	for _, e := range mmForID.ForIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.m1, e.results.err
		}
	}

	if mmForID.ForIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmForID.ForIDMock.defaultExpectation.Counter, 1)
		mm_want := mmForID.ForIDMock.defaultExpectation.params
		mm_got := RecordStorageMockForIDParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmForID.t.Errorf("RecordStorageMock.ForID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmForID.ForIDMock.defaultExpectation.results
		if mm_results == nil {
			mmForID.t.Fatal("No results are set for the RecordStorageMock.ForID")
		}
		return (*mm_results).m1, (*mm_results).err
	}
	if mmForID.funcForID != nil {
		return mmForID.funcForID(ctx, id)
	}
	mmForID.t.Fatalf("Unexpected call to RecordStorageMock.ForID. %v %v", ctx, id)
	return
}

// ForIDAfterCounter returns a count of finished RecordStorageMock.ForID invocations
func (mmForID *RecordStorageMock) ForIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForID.afterForIDCounter)
}

// ForIDBeforeCounter returns a count of RecordStorageMock.ForID invocations
func (mmForID *RecordStorageMock) ForIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForID.beforeForIDCounter)
}

// Calls returns a list of arguments used in each call to RecordStorageMock.ForID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForID *mRecordStorageMockForID) Calls() []*RecordStorageMockForIDParams {
	mmForID.mutex.RLock()

	argCopy := make([]*RecordStorageMockForIDParams, len(mmForID.callArgs))
	copy(argCopy, mmForID.callArgs)

	mmForID.mutex.RUnlock()

	return argCopy
}

// MinimockForIDDone returns true if the count of the ForID invocations corresponds
// the number of defined expectations
func (m *RecordStorageMock) MinimockForIDDone() bool {
	for _, e := range m.ForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForID != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockForIDInspect logs each unmet expectation
func (m *RecordStorageMock) MinimockForIDInspect() {
	for _, e := range m.ForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RecordStorageMock.ForID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		if m.ForIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RecordStorageMock.ForID")
		} else {
			m.t.Errorf("Expected call to RecordStorageMock.ForID with params: %#v", *m.ForIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForID != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		m.t.Error("Expected call to RecordStorageMock.ForID")
	}
}

type mRecordStorageMockSet struct {
	mock               *RecordStorageMock
	defaultExpectation *RecordStorageMockSetExpectation
	expectations       []*RecordStorageMockSetExpectation

	callArgs []*RecordStorageMockSetParams
	mutex    sync.RWMutex
}

// RecordStorageMockSetExpectation specifies expectation struct of the RecordStorage.Set
type RecordStorageMockSetExpectation struct {
	mock    *RecordStorageMock
	params  *RecordStorageMockSetParams
	results *RecordStorageMockSetResults
	Counter uint64
}

// RecordStorageMockSetParams contains parameters of the RecordStorage.Set
type RecordStorageMockSetParams struct {
	ctx context.Context
	rec record.Material
}

// RecordStorageMockSetResults contains results of the RecordStorage.Set
type RecordStorageMockSetResults struct {
	err error
}

// Expect sets up expected params for RecordStorage.Set
func (mmSet *mRecordStorageMockSet) Expect(ctx context.Context, rec record.Material) *mRecordStorageMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RecordStorageMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RecordStorageMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &RecordStorageMockSetParams{ctx, rec}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the RecordStorage.Set
func (mmSet *mRecordStorageMockSet) Inspect(f func(ctx context.Context, rec record.Material)) *mRecordStorageMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for RecordStorageMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by RecordStorage.Set
func (mmSet *mRecordStorageMockSet) Return(err error) *RecordStorageMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RecordStorageMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RecordStorageMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &RecordStorageMockSetResults{err}
	return mmSet.mock
}

//Set uses given function f to mock the RecordStorage.Set method
func (mmSet *mRecordStorageMockSet) Set(f func(ctx context.Context, rec record.Material) (err error)) *RecordStorageMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the RecordStorage.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the RecordStorage.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the RecordStorage.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mRecordStorageMockSet) When(ctx context.Context, rec record.Material) *RecordStorageMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RecordStorageMock.Set mock is already set by Set")
	}

	expectation := &RecordStorageMockSetExpectation{
		mock:   mmSet.mock,
		params: &RecordStorageMockSetParams{ctx, rec},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up RecordStorage.Set return parameters for the expectation previously defined by the When method
func (e *RecordStorageMockSetExpectation) Then(err error) *RecordStorageMock {
	e.results = &RecordStorageMockSetResults{err}
	return e.mock
}

// Set implements RecordStorage
func (mmSet *RecordStorageMock) Set(ctx context.Context, rec record.Material) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(ctx, rec)
	}

	mm_params := &RecordStorageMockSetParams{ctx, rec}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_got := RecordStorageMockSetParams{ctx, rec}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("RecordStorageMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the RecordStorageMock.Set")
		}
		return (*mm_results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(ctx, rec)
	}
	mmSet.t.Fatalf("Unexpected call to RecordStorageMock.Set. %v %v", ctx, rec)
	return
}

// SetAfterCounter returns a count of finished RecordStorageMock.Set invocations
func (mmSet *RecordStorageMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of RecordStorageMock.Set invocations
func (mmSet *RecordStorageMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to RecordStorageMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mRecordStorageMockSet) Calls() []*RecordStorageMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*RecordStorageMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *RecordStorageMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *RecordStorageMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RecordStorageMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RecordStorageMock.Set")
		} else {
			m.t.Errorf("Expected call to RecordStorageMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to RecordStorageMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RecordStorageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockBatchSetInspect()

		m.MinimockForIDInspect()

		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RecordStorageMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RecordStorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockBatchSetDone() &&
		m.MinimockForIDDone() &&
		m.MinimockSetDone()
}
