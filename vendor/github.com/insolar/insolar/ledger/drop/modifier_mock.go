package drop

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ModifierMock implements Modifier
type ModifierMock struct {
	t minimock.Tester

	funcSet          func(ctx context.Context, drop Drop) (err error)
	inspectFuncSet   func(ctx context.Context, drop Drop)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mModifierMockSet
}

// NewModifierMock returns a mock for Modifier
func NewModifierMock(t minimock.Tester) *ModifierMock {
	m := &ModifierMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SetMock = mModifierMockSet{mock: m}
	m.SetMock.callArgs = []*ModifierMockSetParams{}

	return m
}

type mModifierMockSet struct {
	mock               *ModifierMock
	defaultExpectation *ModifierMockSetExpectation
	expectations       []*ModifierMockSetExpectation

	callArgs []*ModifierMockSetParams
	mutex    sync.RWMutex
}

// ModifierMockSetExpectation specifies expectation struct of the Modifier.Set
type ModifierMockSetExpectation struct {
	mock    *ModifierMock
	params  *ModifierMockSetParams
	results *ModifierMockSetResults
	Counter uint64
}

// ModifierMockSetParams contains parameters of the Modifier.Set
type ModifierMockSetParams struct {
	ctx  context.Context
	drop Drop
}

// ModifierMockSetResults contains results of the Modifier.Set
type ModifierMockSetResults struct {
	err error
}

// Expect sets up expected params for Modifier.Set
func (mmSet *mModifierMockSet) Expect(ctx context.Context, drop Drop) *mModifierMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("ModifierMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &ModifierMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &ModifierMockSetParams{ctx, drop}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the Modifier.Set
func (mmSet *mModifierMockSet) Inspect(f func(ctx context.Context, drop Drop)) *mModifierMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for ModifierMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by Modifier.Set
func (mmSet *mModifierMockSet) Return(err error) *ModifierMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("ModifierMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &ModifierMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &ModifierMockSetResults{err}
	return mmSet.mock
}

//Set uses given function f to mock the Modifier.Set method
func (mmSet *mModifierMockSet) Set(f func(ctx context.Context, drop Drop) (err error)) *ModifierMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the Modifier.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the Modifier.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the Modifier.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mModifierMockSet) When(ctx context.Context, drop Drop) *ModifierMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("ModifierMock.Set mock is already set by Set")
	}

	expectation := &ModifierMockSetExpectation{
		mock:   mmSet.mock,
		params: &ModifierMockSetParams{ctx, drop},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up Modifier.Set return parameters for the expectation previously defined by the When method
func (e *ModifierMockSetExpectation) Then(err error) *ModifierMock {
	e.results = &ModifierMockSetResults{err}
	return e.mock
}

// Set implements Modifier
func (mmSet *ModifierMock) Set(ctx context.Context, drop Drop) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(ctx, drop)
	}

	mm_params := &ModifierMockSetParams{ctx, drop}

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
		mm_got := ModifierMockSetParams{ctx, drop}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("ModifierMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the ModifierMock.Set")
		}
		return (*mm_results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(ctx, drop)
	}
	mmSet.t.Fatalf("Unexpected call to ModifierMock.Set. %v %v", ctx, drop)
	return
}

// SetAfterCounter returns a count of finished ModifierMock.Set invocations
func (mmSet *ModifierMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of ModifierMock.Set invocations
func (mmSet *ModifierMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to ModifierMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mModifierMockSet) Calls() []*ModifierMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*ModifierMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *ModifierMock) MinimockSetDone() bool {
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
func (m *ModifierMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ModifierMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ModifierMock.Set")
		} else {
			m.t.Errorf("Expected call to ModifierMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to ModifierMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ModifierMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ModifierMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ModifierMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSetDone()
}