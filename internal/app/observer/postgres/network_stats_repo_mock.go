package postgres

// Code generated by http://github.com/gojuno/minimock (2.1.9). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// NetworkStatsRepoMock implements NetworkStatsRepo
type NetworkStatsRepoMock struct {
	t minimock.Tester

	funcCountStats          func() (n1 NetworkStatsModel, err error)
	inspectFuncCountStats   func()
	afterCountStatsCounter  uint64
	beforeCountStatsCounter uint64
	CountStatsMock          mNetworkStatsRepoMockCountStats

	funcInsertStats          func(n1 NetworkStatsModel) (err error)
	inspectFuncInsertStats   func(n1 NetworkStatsModel)
	afterInsertStatsCounter  uint64
	beforeInsertStatsCounter uint64
	InsertStatsMock          mNetworkStatsRepoMockInsertStats

	funcLastStats          func() (n1 NetworkStatsModel, err error)
	inspectFuncLastStats   func()
	afterLastStatsCounter  uint64
	beforeLastStatsCounter uint64
	LastStatsMock          mNetworkStatsRepoMockLastStats
}

// NewNetworkStatsRepoMock returns a mock for NetworkStatsRepo
func NewNetworkStatsRepoMock(t minimock.Tester) *NetworkStatsRepoMock {
	m := &NetworkStatsRepoMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CountStatsMock = mNetworkStatsRepoMockCountStats{mock: m}

	m.InsertStatsMock = mNetworkStatsRepoMockInsertStats{mock: m}
	m.InsertStatsMock.callArgs = []*NetworkStatsRepoMockInsertStatsParams{}

	m.LastStatsMock = mNetworkStatsRepoMockLastStats{mock: m}

	return m
}

type mNetworkStatsRepoMockCountStats struct {
	mock               *NetworkStatsRepoMock
	defaultExpectation *NetworkStatsRepoMockCountStatsExpectation
	expectations       []*NetworkStatsRepoMockCountStatsExpectation
}

// NetworkStatsRepoMockCountStatsExpectation specifies expectation struct of the NetworkStatsRepo.CountStats
type NetworkStatsRepoMockCountStatsExpectation struct {
	mock *NetworkStatsRepoMock

	results *NetworkStatsRepoMockCountStatsResults
	Counter uint64
}

// NetworkStatsRepoMockCountStatsResults contains results of the NetworkStatsRepo.CountStats
type NetworkStatsRepoMockCountStatsResults struct {
	n1  NetworkStatsModel
	err error
}

// Expect sets up expected params for NetworkStatsRepo.CountStats
func (mmCountStats *mNetworkStatsRepoMockCountStats) Expect() *mNetworkStatsRepoMockCountStats {
	if mmCountStats.mock.funcCountStats != nil {
		mmCountStats.mock.t.Fatalf("NetworkStatsRepoMock.CountStats mock is already set by Set")
	}

	if mmCountStats.defaultExpectation == nil {
		mmCountStats.defaultExpectation = &NetworkStatsRepoMockCountStatsExpectation{}
	}

	return mmCountStats
}

// Inspect accepts an inspector function that has same arguments as the NetworkStatsRepo.CountStats
func (mmCountStats *mNetworkStatsRepoMockCountStats) Inspect(f func()) *mNetworkStatsRepoMockCountStats {
	if mmCountStats.mock.inspectFuncCountStats != nil {
		mmCountStats.mock.t.Fatalf("Inspect function is already set for NetworkStatsRepoMock.CountStats")
	}

	mmCountStats.mock.inspectFuncCountStats = f

	return mmCountStats
}

// Return sets up results that will be returned by NetworkStatsRepo.CountStats
func (mmCountStats *mNetworkStatsRepoMockCountStats) Return(n1 NetworkStatsModel, err error) *NetworkStatsRepoMock {
	if mmCountStats.mock.funcCountStats != nil {
		mmCountStats.mock.t.Fatalf("NetworkStatsRepoMock.CountStats mock is already set by Set")
	}

	if mmCountStats.defaultExpectation == nil {
		mmCountStats.defaultExpectation = &NetworkStatsRepoMockCountStatsExpectation{mock: mmCountStats.mock}
	}
	mmCountStats.defaultExpectation.results = &NetworkStatsRepoMockCountStatsResults{n1, err}
	return mmCountStats.mock
}

//Set uses given function f to mock the NetworkStatsRepo.CountStats method
func (mmCountStats *mNetworkStatsRepoMockCountStats) Set(f func() (n1 NetworkStatsModel, err error)) *NetworkStatsRepoMock {
	if mmCountStats.defaultExpectation != nil {
		mmCountStats.mock.t.Fatalf("Default expectation is already set for the NetworkStatsRepo.CountStats method")
	}

	if len(mmCountStats.expectations) > 0 {
		mmCountStats.mock.t.Fatalf("Some expectations are already set for the NetworkStatsRepo.CountStats method")
	}

	mmCountStats.mock.funcCountStats = f
	return mmCountStats.mock
}

// CountStats implements NetworkStatsRepo
func (mmCountStats *NetworkStatsRepoMock) CountStats() (n1 NetworkStatsModel, err error) {
	mm_atomic.AddUint64(&mmCountStats.beforeCountStatsCounter, 1)
	defer mm_atomic.AddUint64(&mmCountStats.afterCountStatsCounter, 1)

	if mmCountStats.inspectFuncCountStats != nil {
		mmCountStats.inspectFuncCountStats()
	}

	if mmCountStats.CountStatsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCountStats.CountStatsMock.defaultExpectation.Counter, 1)

		mm_results := mmCountStats.CountStatsMock.defaultExpectation.results
		if mm_results == nil {
			mmCountStats.t.Fatal("No results are set for the NetworkStatsRepoMock.CountStats")
		}
		return (*mm_results).n1, (*mm_results).err
	}
	if mmCountStats.funcCountStats != nil {
		return mmCountStats.funcCountStats()
	}
	mmCountStats.t.Fatalf("Unexpected call to NetworkStatsRepoMock.CountStats.")
	return
}

// CountStatsAfterCounter returns a count of finished NetworkStatsRepoMock.CountStats invocations
func (mmCountStats *NetworkStatsRepoMock) CountStatsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCountStats.afterCountStatsCounter)
}

// CountStatsBeforeCounter returns a count of NetworkStatsRepoMock.CountStats invocations
func (mmCountStats *NetworkStatsRepoMock) CountStatsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCountStats.beforeCountStatsCounter)
}

// MinimockCountStatsDone returns true if the count of the CountStats invocations corresponds
// the number of defined expectations
func (m *NetworkStatsRepoMock) MinimockCountStatsDone() bool {
	for _, e := range m.CountStatsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CountStatsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCountStatsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCountStats != nil && mm_atomic.LoadUint64(&m.afterCountStatsCounter) < 1 {
		return false
	}
	return true
}

// MinimockCountStatsInspect logs each unmet expectation
func (m *NetworkStatsRepoMock) MinimockCountStatsInspect() {
	for _, e := range m.CountStatsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to NetworkStatsRepoMock.CountStats")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CountStatsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCountStatsCounter) < 1 {
		m.t.Error("Expected call to NetworkStatsRepoMock.CountStats")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCountStats != nil && mm_atomic.LoadUint64(&m.afterCountStatsCounter) < 1 {
		m.t.Error("Expected call to NetworkStatsRepoMock.CountStats")
	}
}

type mNetworkStatsRepoMockInsertStats struct {
	mock               *NetworkStatsRepoMock
	defaultExpectation *NetworkStatsRepoMockInsertStatsExpectation
	expectations       []*NetworkStatsRepoMockInsertStatsExpectation

	callArgs []*NetworkStatsRepoMockInsertStatsParams
	mutex    sync.RWMutex
}

// NetworkStatsRepoMockInsertStatsExpectation specifies expectation struct of the NetworkStatsRepo.InsertStats
type NetworkStatsRepoMockInsertStatsExpectation struct {
	mock    *NetworkStatsRepoMock
	params  *NetworkStatsRepoMockInsertStatsParams
	results *NetworkStatsRepoMockInsertStatsResults
	Counter uint64
}

// NetworkStatsRepoMockInsertStatsParams contains parameters of the NetworkStatsRepo.InsertStats
type NetworkStatsRepoMockInsertStatsParams struct {
	n1 NetworkStatsModel
}

// NetworkStatsRepoMockInsertStatsResults contains results of the NetworkStatsRepo.InsertStats
type NetworkStatsRepoMockInsertStatsResults struct {
	err error
}

// Expect sets up expected params for NetworkStatsRepo.InsertStats
func (mmInsertStats *mNetworkStatsRepoMockInsertStats) Expect(n1 NetworkStatsModel) *mNetworkStatsRepoMockInsertStats {
	if mmInsertStats.mock.funcInsertStats != nil {
		mmInsertStats.mock.t.Fatalf("NetworkStatsRepoMock.InsertStats mock is already set by Set")
	}

	if mmInsertStats.defaultExpectation == nil {
		mmInsertStats.defaultExpectation = &NetworkStatsRepoMockInsertStatsExpectation{}
	}

	mmInsertStats.defaultExpectation.params = &NetworkStatsRepoMockInsertStatsParams{n1}
	for _, e := range mmInsertStats.expectations {
		if minimock.Equal(e.params, mmInsertStats.defaultExpectation.params) {
			mmInsertStats.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInsertStats.defaultExpectation.params)
		}
	}

	return mmInsertStats
}

// Inspect accepts an inspector function that has same arguments as the NetworkStatsRepo.InsertStats
func (mmInsertStats *mNetworkStatsRepoMockInsertStats) Inspect(f func(n1 NetworkStatsModel)) *mNetworkStatsRepoMockInsertStats {
	if mmInsertStats.mock.inspectFuncInsertStats != nil {
		mmInsertStats.mock.t.Fatalf("Inspect function is already set for NetworkStatsRepoMock.InsertStats")
	}

	mmInsertStats.mock.inspectFuncInsertStats = f

	return mmInsertStats
}

// Return sets up results that will be returned by NetworkStatsRepo.InsertStats
func (mmInsertStats *mNetworkStatsRepoMockInsertStats) Return(err error) *NetworkStatsRepoMock {
	if mmInsertStats.mock.funcInsertStats != nil {
		mmInsertStats.mock.t.Fatalf("NetworkStatsRepoMock.InsertStats mock is already set by Set")
	}

	if mmInsertStats.defaultExpectation == nil {
		mmInsertStats.defaultExpectation = &NetworkStatsRepoMockInsertStatsExpectation{mock: mmInsertStats.mock}
	}
	mmInsertStats.defaultExpectation.results = &NetworkStatsRepoMockInsertStatsResults{err}
	return mmInsertStats.mock
}

//Set uses given function f to mock the NetworkStatsRepo.InsertStats method
func (mmInsertStats *mNetworkStatsRepoMockInsertStats) Set(f func(n1 NetworkStatsModel) (err error)) *NetworkStatsRepoMock {
	if mmInsertStats.defaultExpectation != nil {
		mmInsertStats.mock.t.Fatalf("Default expectation is already set for the NetworkStatsRepo.InsertStats method")
	}

	if len(mmInsertStats.expectations) > 0 {
		mmInsertStats.mock.t.Fatalf("Some expectations are already set for the NetworkStatsRepo.InsertStats method")
	}

	mmInsertStats.mock.funcInsertStats = f
	return mmInsertStats.mock
}

// When sets expectation for the NetworkStatsRepo.InsertStats which will trigger the result defined by the following
// Then helper
func (mmInsertStats *mNetworkStatsRepoMockInsertStats) When(n1 NetworkStatsModel) *NetworkStatsRepoMockInsertStatsExpectation {
	if mmInsertStats.mock.funcInsertStats != nil {
		mmInsertStats.mock.t.Fatalf("NetworkStatsRepoMock.InsertStats mock is already set by Set")
	}

	expectation := &NetworkStatsRepoMockInsertStatsExpectation{
		mock:   mmInsertStats.mock,
		params: &NetworkStatsRepoMockInsertStatsParams{n1},
	}
	mmInsertStats.expectations = append(mmInsertStats.expectations, expectation)
	return expectation
}

// Then sets up NetworkStatsRepo.InsertStats return parameters for the expectation previously defined by the When method
func (e *NetworkStatsRepoMockInsertStatsExpectation) Then(err error) *NetworkStatsRepoMock {
	e.results = &NetworkStatsRepoMockInsertStatsResults{err}
	return e.mock
}

// InsertStats implements NetworkStatsRepo
func (mmInsertStats *NetworkStatsRepoMock) InsertStats(n1 NetworkStatsModel) (err error) {
	mm_atomic.AddUint64(&mmInsertStats.beforeInsertStatsCounter, 1)
	defer mm_atomic.AddUint64(&mmInsertStats.afterInsertStatsCounter, 1)

	if mmInsertStats.inspectFuncInsertStats != nil {
		mmInsertStats.inspectFuncInsertStats(n1)
	}

	mm_params := &NetworkStatsRepoMockInsertStatsParams{n1}

	// Record call args
	mmInsertStats.InsertStatsMock.mutex.Lock()
	mmInsertStats.InsertStatsMock.callArgs = append(mmInsertStats.InsertStatsMock.callArgs, mm_params)
	mmInsertStats.InsertStatsMock.mutex.Unlock()

	for _, e := range mmInsertStats.InsertStatsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmInsertStats.InsertStatsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInsertStats.InsertStatsMock.defaultExpectation.Counter, 1)
		mm_want := mmInsertStats.InsertStatsMock.defaultExpectation.params
		mm_got := NetworkStatsRepoMockInsertStatsParams{n1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInsertStats.t.Errorf("NetworkStatsRepoMock.InsertStats got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInsertStats.InsertStatsMock.defaultExpectation.results
		if mm_results == nil {
			mmInsertStats.t.Fatal("No results are set for the NetworkStatsRepoMock.InsertStats")
		}
		return (*mm_results).err
	}
	if mmInsertStats.funcInsertStats != nil {
		return mmInsertStats.funcInsertStats(n1)
	}
	mmInsertStats.t.Fatalf("Unexpected call to NetworkStatsRepoMock.InsertStats. %v", n1)
	return
}

// InsertStatsAfterCounter returns a count of finished NetworkStatsRepoMock.InsertStats invocations
func (mmInsertStats *NetworkStatsRepoMock) InsertStatsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsertStats.afterInsertStatsCounter)
}

// InsertStatsBeforeCounter returns a count of NetworkStatsRepoMock.InsertStats invocations
func (mmInsertStats *NetworkStatsRepoMock) InsertStatsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInsertStats.beforeInsertStatsCounter)
}

// Calls returns a list of arguments used in each call to NetworkStatsRepoMock.InsertStats.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInsertStats *mNetworkStatsRepoMockInsertStats) Calls() []*NetworkStatsRepoMockInsertStatsParams {
	mmInsertStats.mutex.RLock()

	argCopy := make([]*NetworkStatsRepoMockInsertStatsParams, len(mmInsertStats.callArgs))
	copy(argCopy, mmInsertStats.callArgs)

	mmInsertStats.mutex.RUnlock()

	return argCopy
}

// MinimockInsertStatsDone returns true if the count of the InsertStats invocations corresponds
// the number of defined expectations
func (m *NetworkStatsRepoMock) MinimockInsertStatsDone() bool {
	for _, e := range m.InsertStatsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertStatsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInsertStatsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertStats != nil && mm_atomic.LoadUint64(&m.afterInsertStatsCounter) < 1 {
		return false
	}
	return true
}

// MinimockInsertStatsInspect logs each unmet expectation
func (m *NetworkStatsRepoMock) MinimockInsertStatsInspect() {
	for _, e := range m.InsertStatsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NetworkStatsRepoMock.InsertStats with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InsertStatsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInsertStatsCounter) < 1 {
		if m.InsertStatsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to NetworkStatsRepoMock.InsertStats")
		} else {
			m.t.Errorf("Expected call to NetworkStatsRepoMock.InsertStats with params: %#v", *m.InsertStatsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInsertStats != nil && mm_atomic.LoadUint64(&m.afterInsertStatsCounter) < 1 {
		m.t.Error("Expected call to NetworkStatsRepoMock.InsertStats")
	}
}

type mNetworkStatsRepoMockLastStats struct {
	mock               *NetworkStatsRepoMock
	defaultExpectation *NetworkStatsRepoMockLastStatsExpectation
	expectations       []*NetworkStatsRepoMockLastStatsExpectation
}

// NetworkStatsRepoMockLastStatsExpectation specifies expectation struct of the NetworkStatsRepo.LastStats
type NetworkStatsRepoMockLastStatsExpectation struct {
	mock *NetworkStatsRepoMock

	results *NetworkStatsRepoMockLastStatsResults
	Counter uint64
}

// NetworkStatsRepoMockLastStatsResults contains results of the NetworkStatsRepo.LastStats
type NetworkStatsRepoMockLastStatsResults struct {
	n1  NetworkStatsModel
	err error
}

// Expect sets up expected params for NetworkStatsRepo.LastStats
func (mmLastStats *mNetworkStatsRepoMockLastStats) Expect() *mNetworkStatsRepoMockLastStats {
	if mmLastStats.mock.funcLastStats != nil {
		mmLastStats.mock.t.Fatalf("NetworkStatsRepoMock.LastStats mock is already set by Set")
	}

	if mmLastStats.defaultExpectation == nil {
		mmLastStats.defaultExpectation = &NetworkStatsRepoMockLastStatsExpectation{}
	}

	return mmLastStats
}

// Inspect accepts an inspector function that has same arguments as the NetworkStatsRepo.LastStats
func (mmLastStats *mNetworkStatsRepoMockLastStats) Inspect(f func()) *mNetworkStatsRepoMockLastStats {
	if mmLastStats.mock.inspectFuncLastStats != nil {
		mmLastStats.mock.t.Fatalf("Inspect function is already set for NetworkStatsRepoMock.LastStats")
	}

	mmLastStats.mock.inspectFuncLastStats = f

	return mmLastStats
}

// Return sets up results that will be returned by NetworkStatsRepo.LastStats
func (mmLastStats *mNetworkStatsRepoMockLastStats) Return(n1 NetworkStatsModel, err error) *NetworkStatsRepoMock {
	if mmLastStats.mock.funcLastStats != nil {
		mmLastStats.mock.t.Fatalf("NetworkStatsRepoMock.LastStats mock is already set by Set")
	}

	if mmLastStats.defaultExpectation == nil {
		mmLastStats.defaultExpectation = &NetworkStatsRepoMockLastStatsExpectation{mock: mmLastStats.mock}
	}
	mmLastStats.defaultExpectation.results = &NetworkStatsRepoMockLastStatsResults{n1, err}
	return mmLastStats.mock
}

//Set uses given function f to mock the NetworkStatsRepo.LastStats method
func (mmLastStats *mNetworkStatsRepoMockLastStats) Set(f func() (n1 NetworkStatsModel, err error)) *NetworkStatsRepoMock {
	if mmLastStats.defaultExpectation != nil {
		mmLastStats.mock.t.Fatalf("Default expectation is already set for the NetworkStatsRepo.LastStats method")
	}

	if len(mmLastStats.expectations) > 0 {
		mmLastStats.mock.t.Fatalf("Some expectations are already set for the NetworkStatsRepo.LastStats method")
	}

	mmLastStats.mock.funcLastStats = f
	return mmLastStats.mock
}

// LastStats implements NetworkStatsRepo
func (mmLastStats *NetworkStatsRepoMock) LastStats() (n1 NetworkStatsModel, err error) {
	mm_atomic.AddUint64(&mmLastStats.beforeLastStatsCounter, 1)
	defer mm_atomic.AddUint64(&mmLastStats.afterLastStatsCounter, 1)

	if mmLastStats.inspectFuncLastStats != nil {
		mmLastStats.inspectFuncLastStats()
	}

	if mmLastStats.LastStatsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLastStats.LastStatsMock.defaultExpectation.Counter, 1)

		mm_results := mmLastStats.LastStatsMock.defaultExpectation.results
		if mm_results == nil {
			mmLastStats.t.Fatal("No results are set for the NetworkStatsRepoMock.LastStats")
		}
		return (*mm_results).n1, (*mm_results).err
	}
	if mmLastStats.funcLastStats != nil {
		return mmLastStats.funcLastStats()
	}
	mmLastStats.t.Fatalf("Unexpected call to NetworkStatsRepoMock.LastStats.")
	return
}

// LastStatsAfterCounter returns a count of finished NetworkStatsRepoMock.LastStats invocations
func (mmLastStats *NetworkStatsRepoMock) LastStatsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLastStats.afterLastStatsCounter)
}

// LastStatsBeforeCounter returns a count of NetworkStatsRepoMock.LastStats invocations
func (mmLastStats *NetworkStatsRepoMock) LastStatsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLastStats.beforeLastStatsCounter)
}

// MinimockLastStatsDone returns true if the count of the LastStats invocations corresponds
// the number of defined expectations
func (m *NetworkStatsRepoMock) MinimockLastStatsDone() bool {
	for _, e := range m.LastStatsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LastStatsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLastStatsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLastStats != nil && mm_atomic.LoadUint64(&m.afterLastStatsCounter) < 1 {
		return false
	}
	return true
}

// MinimockLastStatsInspect logs each unmet expectation
func (m *NetworkStatsRepoMock) MinimockLastStatsInspect() {
	for _, e := range m.LastStatsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to NetworkStatsRepoMock.LastStats")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LastStatsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLastStatsCounter) < 1 {
		m.t.Error("Expected call to NetworkStatsRepoMock.LastStats")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLastStats != nil && mm_atomic.LoadUint64(&m.afterLastStatsCounter) < 1 {
		m.t.Error("Expected call to NetworkStatsRepoMock.LastStats")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *NetworkStatsRepoMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCountStatsInspect()

		m.MinimockInsertStatsInspect()

		m.MinimockLastStatsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *NetworkStatsRepoMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *NetworkStatsRepoMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCountStatsDone() &&
		m.MinimockInsertStatsDone() &&
		m.MinimockLastStatsDone()
}
