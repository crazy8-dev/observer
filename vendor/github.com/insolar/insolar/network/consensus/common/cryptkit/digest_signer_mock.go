package cryptkit

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// DigestSignerMock implements DigestSigner
type DigestSignerMock struct {
	t minimock.Tester

	funcGetSignMethod          func() (s1 SignMethod)
	inspectFuncGetSignMethod   func()
	afterGetSignMethodCounter  uint64
	beforeGetSignMethodCounter uint64
	GetSignMethodMock          mDigestSignerMockGetSignMethod

	funcSignDigest          func(digest Digest) (s1 Signature)
	inspectFuncSignDigest   func(digest Digest)
	afterSignDigestCounter  uint64
	beforeSignDigestCounter uint64
	SignDigestMock          mDigestSignerMockSignDigest
}

// NewDigestSignerMock returns a mock for DigestSigner
func NewDigestSignerMock(t minimock.Tester) *DigestSignerMock {
	m := &DigestSignerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetSignMethodMock = mDigestSignerMockGetSignMethod{mock: m}

	m.SignDigestMock = mDigestSignerMockSignDigest{mock: m}
	m.SignDigestMock.callArgs = []*DigestSignerMockSignDigestParams{}

	return m
}

type mDigestSignerMockGetSignMethod struct {
	mock               *DigestSignerMock
	defaultExpectation *DigestSignerMockGetSignMethodExpectation
	expectations       []*DigestSignerMockGetSignMethodExpectation
}

// DigestSignerMockGetSignMethodExpectation specifies expectation struct of the DigestSigner.GetSignMethod
type DigestSignerMockGetSignMethodExpectation struct {
	mock *DigestSignerMock

	results *DigestSignerMockGetSignMethodResults
	Counter uint64
}

// DigestSignerMockGetSignMethodResults contains results of the DigestSigner.GetSignMethod
type DigestSignerMockGetSignMethodResults struct {
	s1 SignMethod
}

// Expect sets up expected params for DigestSigner.GetSignMethod
func (mmGetSignMethod *mDigestSignerMockGetSignMethod) Expect() *mDigestSignerMockGetSignMethod {
	if mmGetSignMethod.mock.funcGetSignMethod != nil {
		mmGetSignMethod.mock.t.Fatalf("DigestSignerMock.GetSignMethod mock is already set by Set")
	}

	if mmGetSignMethod.defaultExpectation == nil {
		mmGetSignMethod.defaultExpectation = &DigestSignerMockGetSignMethodExpectation{}
	}

	return mmGetSignMethod
}

// Inspect accepts an inspector function that has same arguments as the DigestSigner.GetSignMethod
func (mmGetSignMethod *mDigestSignerMockGetSignMethod) Inspect(f func()) *mDigestSignerMockGetSignMethod {
	if mmGetSignMethod.mock.inspectFuncGetSignMethod != nil {
		mmGetSignMethod.mock.t.Fatalf("Inspect function is already set for DigestSignerMock.GetSignMethod")
	}

	mmGetSignMethod.mock.inspectFuncGetSignMethod = f

	return mmGetSignMethod
}

// Return sets up results that will be returned by DigestSigner.GetSignMethod
func (mmGetSignMethod *mDigestSignerMockGetSignMethod) Return(s1 SignMethod) *DigestSignerMock {
	if mmGetSignMethod.mock.funcGetSignMethod != nil {
		mmGetSignMethod.mock.t.Fatalf("DigestSignerMock.GetSignMethod mock is already set by Set")
	}

	if mmGetSignMethod.defaultExpectation == nil {
		mmGetSignMethod.defaultExpectation = &DigestSignerMockGetSignMethodExpectation{mock: mmGetSignMethod.mock}
	}
	mmGetSignMethod.defaultExpectation.results = &DigestSignerMockGetSignMethodResults{s1}
	return mmGetSignMethod.mock
}

//Set uses given function f to mock the DigestSigner.GetSignMethod method
func (mmGetSignMethod *mDigestSignerMockGetSignMethod) Set(f func() (s1 SignMethod)) *DigestSignerMock {
	if mmGetSignMethod.defaultExpectation != nil {
		mmGetSignMethod.mock.t.Fatalf("Default expectation is already set for the DigestSigner.GetSignMethod method")
	}

	if len(mmGetSignMethod.expectations) > 0 {
		mmGetSignMethod.mock.t.Fatalf("Some expectations are already set for the DigestSigner.GetSignMethod method")
	}

	mmGetSignMethod.mock.funcGetSignMethod = f
	return mmGetSignMethod.mock
}

// GetSignMethod implements DigestSigner
func (mmGetSignMethod *DigestSignerMock) GetSignMethod() (s1 SignMethod) {
	mm_atomic.AddUint64(&mmGetSignMethod.beforeGetSignMethodCounter, 1)
	defer mm_atomic.AddUint64(&mmGetSignMethod.afterGetSignMethodCounter, 1)

	if mmGetSignMethod.inspectFuncGetSignMethod != nil {
		mmGetSignMethod.inspectFuncGetSignMethod()
	}

	if mmGetSignMethod.GetSignMethodMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetSignMethod.GetSignMethodMock.defaultExpectation.Counter, 1)

		results := mmGetSignMethod.GetSignMethodMock.defaultExpectation.results
		if results == nil {
			mmGetSignMethod.t.Fatal("No results are set for the DigestSignerMock.GetSignMethod")
		}
		return (*results).s1
	}
	if mmGetSignMethod.funcGetSignMethod != nil {
		return mmGetSignMethod.funcGetSignMethod()
	}
	mmGetSignMethod.t.Fatalf("Unexpected call to DigestSignerMock.GetSignMethod.")
	return
}

// GetSignMethodAfterCounter returns a count of finished DigestSignerMock.GetSignMethod invocations
func (mmGetSignMethod *DigestSignerMock) GetSignMethodAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSignMethod.afterGetSignMethodCounter)
}

// GetSignMethodBeforeCounter returns a count of DigestSignerMock.GetSignMethod invocations
func (mmGetSignMethod *DigestSignerMock) GetSignMethodBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSignMethod.beforeGetSignMethodCounter)
}

// MinimockGetSignMethodDone returns true if the count of the GetSignMethod invocations corresponds
// the number of defined expectations
func (m *DigestSignerMock) MinimockGetSignMethodDone() bool {
	for _, e := range m.GetSignMethodMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSignMethodMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSignMethodCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSignMethod != nil && mm_atomic.LoadUint64(&m.afterGetSignMethodCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetSignMethodInspect logs each unmet expectation
func (m *DigestSignerMock) MinimockGetSignMethodInspect() {
	for _, e := range m.GetSignMethodMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to DigestSignerMock.GetSignMethod")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSignMethodMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSignMethodCounter) < 1 {
		m.t.Error("Expected call to DigestSignerMock.GetSignMethod")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSignMethod != nil && mm_atomic.LoadUint64(&m.afterGetSignMethodCounter) < 1 {
		m.t.Error("Expected call to DigestSignerMock.GetSignMethod")
	}
}

type mDigestSignerMockSignDigest struct {
	mock               *DigestSignerMock
	defaultExpectation *DigestSignerMockSignDigestExpectation
	expectations       []*DigestSignerMockSignDigestExpectation

	callArgs []*DigestSignerMockSignDigestParams
	mutex    sync.RWMutex
}

// DigestSignerMockSignDigestExpectation specifies expectation struct of the DigestSigner.SignDigest
type DigestSignerMockSignDigestExpectation struct {
	mock    *DigestSignerMock
	params  *DigestSignerMockSignDigestParams
	results *DigestSignerMockSignDigestResults
	Counter uint64
}

// DigestSignerMockSignDigestParams contains parameters of the DigestSigner.SignDigest
type DigestSignerMockSignDigestParams struct {
	digest Digest
}

// DigestSignerMockSignDigestResults contains results of the DigestSigner.SignDigest
type DigestSignerMockSignDigestResults struct {
	s1 Signature
}

// Expect sets up expected params for DigestSigner.SignDigest
func (mmSignDigest *mDigestSignerMockSignDigest) Expect(digest Digest) *mDigestSignerMockSignDigest {
	if mmSignDigest.mock.funcSignDigest != nil {
		mmSignDigest.mock.t.Fatalf("DigestSignerMock.SignDigest mock is already set by Set")
	}

	if mmSignDigest.defaultExpectation == nil {
		mmSignDigest.defaultExpectation = &DigestSignerMockSignDigestExpectation{}
	}

	mmSignDigest.defaultExpectation.params = &DigestSignerMockSignDigestParams{digest}
	for _, e := range mmSignDigest.expectations {
		if minimock.Equal(e.params, mmSignDigest.defaultExpectation.params) {
			mmSignDigest.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSignDigest.defaultExpectation.params)
		}
	}

	return mmSignDigest
}

// Inspect accepts an inspector function that has same arguments as the DigestSigner.SignDigest
func (mmSignDigest *mDigestSignerMockSignDigest) Inspect(f func(digest Digest)) *mDigestSignerMockSignDigest {
	if mmSignDigest.mock.inspectFuncSignDigest != nil {
		mmSignDigest.mock.t.Fatalf("Inspect function is already set for DigestSignerMock.SignDigest")
	}

	mmSignDigest.mock.inspectFuncSignDigest = f

	return mmSignDigest
}

// Return sets up results that will be returned by DigestSigner.SignDigest
func (mmSignDigest *mDigestSignerMockSignDigest) Return(s1 Signature) *DigestSignerMock {
	if mmSignDigest.mock.funcSignDigest != nil {
		mmSignDigest.mock.t.Fatalf("DigestSignerMock.SignDigest mock is already set by Set")
	}

	if mmSignDigest.defaultExpectation == nil {
		mmSignDigest.defaultExpectation = &DigestSignerMockSignDigestExpectation{mock: mmSignDigest.mock}
	}
	mmSignDigest.defaultExpectation.results = &DigestSignerMockSignDigestResults{s1}
	return mmSignDigest.mock
}

//Set uses given function f to mock the DigestSigner.SignDigest method
func (mmSignDigest *mDigestSignerMockSignDigest) Set(f func(digest Digest) (s1 Signature)) *DigestSignerMock {
	if mmSignDigest.defaultExpectation != nil {
		mmSignDigest.mock.t.Fatalf("Default expectation is already set for the DigestSigner.SignDigest method")
	}

	if len(mmSignDigest.expectations) > 0 {
		mmSignDigest.mock.t.Fatalf("Some expectations are already set for the DigestSigner.SignDigest method")
	}

	mmSignDigest.mock.funcSignDigest = f
	return mmSignDigest.mock
}

// When sets expectation for the DigestSigner.SignDigest which will trigger the result defined by the following
// Then helper
func (mmSignDigest *mDigestSignerMockSignDigest) When(digest Digest) *DigestSignerMockSignDigestExpectation {
	if mmSignDigest.mock.funcSignDigest != nil {
		mmSignDigest.mock.t.Fatalf("DigestSignerMock.SignDigest mock is already set by Set")
	}

	expectation := &DigestSignerMockSignDigestExpectation{
		mock:   mmSignDigest.mock,
		params: &DigestSignerMockSignDigestParams{digest},
	}
	mmSignDigest.expectations = append(mmSignDigest.expectations, expectation)
	return expectation
}

// Then sets up DigestSigner.SignDigest return parameters for the expectation previously defined by the When method
func (e *DigestSignerMockSignDigestExpectation) Then(s1 Signature) *DigestSignerMock {
	e.results = &DigestSignerMockSignDigestResults{s1}
	return e.mock
}

// SignDigest implements DigestSigner
func (mmSignDigest *DigestSignerMock) SignDigest(digest Digest) (s1 Signature) {
	mm_atomic.AddUint64(&mmSignDigest.beforeSignDigestCounter, 1)
	defer mm_atomic.AddUint64(&mmSignDigest.afterSignDigestCounter, 1)

	if mmSignDigest.inspectFuncSignDigest != nil {
		mmSignDigest.inspectFuncSignDigest(digest)
	}

	params := &DigestSignerMockSignDigestParams{digest}

	// Record call args
	mmSignDigest.SignDigestMock.mutex.Lock()
	mmSignDigest.SignDigestMock.callArgs = append(mmSignDigest.SignDigestMock.callArgs, params)
	mmSignDigest.SignDigestMock.mutex.Unlock()

	for _, e := range mmSignDigest.SignDigestMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1
		}
	}

	if mmSignDigest.SignDigestMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSignDigest.SignDigestMock.defaultExpectation.Counter, 1)
		want := mmSignDigest.SignDigestMock.defaultExpectation.params
		got := DigestSignerMockSignDigestParams{digest}
		if want != nil && !minimock.Equal(*want, got) {
			mmSignDigest.t.Errorf("DigestSignerMock.SignDigest got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSignDigest.SignDigestMock.defaultExpectation.results
		if results == nil {
			mmSignDigest.t.Fatal("No results are set for the DigestSignerMock.SignDigest")
		}
		return (*results).s1
	}
	if mmSignDigest.funcSignDigest != nil {
		return mmSignDigest.funcSignDigest(digest)
	}
	mmSignDigest.t.Fatalf("Unexpected call to DigestSignerMock.SignDigest. %v", digest)
	return
}

// SignDigestAfterCounter returns a count of finished DigestSignerMock.SignDigest invocations
func (mmSignDigest *DigestSignerMock) SignDigestAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSignDigest.afterSignDigestCounter)
}

// SignDigestBeforeCounter returns a count of DigestSignerMock.SignDigest invocations
func (mmSignDigest *DigestSignerMock) SignDigestBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSignDigest.beforeSignDigestCounter)
}

// Calls returns a list of arguments used in each call to DigestSignerMock.SignDigest.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSignDigest *mDigestSignerMockSignDigest) Calls() []*DigestSignerMockSignDigestParams {
	mmSignDigest.mutex.RLock()

	argCopy := make([]*DigestSignerMockSignDigestParams, len(mmSignDigest.callArgs))
	copy(argCopy, mmSignDigest.callArgs)

	mmSignDigest.mutex.RUnlock()

	return argCopy
}

// MinimockSignDigestDone returns true if the count of the SignDigest invocations corresponds
// the number of defined expectations
func (m *DigestSignerMock) MinimockSignDigestDone() bool {
	for _, e := range m.SignDigestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SignDigestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSignDigestCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSignDigest != nil && mm_atomic.LoadUint64(&m.afterSignDigestCounter) < 1 {
		return false
	}
	return true
}

// MinimockSignDigestInspect logs each unmet expectation
func (m *DigestSignerMock) MinimockSignDigestInspect() {
	for _, e := range m.SignDigestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DigestSignerMock.SignDigest with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SignDigestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSignDigestCounter) < 1 {
		if m.SignDigestMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DigestSignerMock.SignDigest")
		} else {
			m.t.Errorf("Expected call to DigestSignerMock.SignDigest with params: %#v", *m.SignDigestMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSignDigest != nil && mm_atomic.LoadUint64(&m.afterSignDigestCounter) < 1 {
		m.t.Error("Expected call to DigestSignerMock.SignDigest")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DigestSignerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetSignMethodInspect()

		m.MinimockSignDigestInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DigestSignerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *DigestSignerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetSignMethodDone() &&
		m.MinimockSignDigestDone()
}