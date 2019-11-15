package artifacts

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/insolar"
)

// CodeDescriptorMock implements CodeDescriptor
type CodeDescriptorMock struct {
	t minimock.Tester

	funcCode          func() (ba1 []byte, err error)
	inspectFuncCode   func()
	afterCodeCounter  uint64
	beforeCodeCounter uint64
	CodeMock          mCodeDescriptorMockCode

	funcMachineType          func() (m1 insolar.MachineType)
	inspectFuncMachineType   func()
	afterMachineTypeCounter  uint64
	beforeMachineTypeCounter uint64
	MachineTypeMock          mCodeDescriptorMockMachineType

	funcRef          func() (rp1 *insolar.Reference)
	inspectFuncRef   func()
	afterRefCounter  uint64
	beforeRefCounter uint64
	RefMock          mCodeDescriptorMockRef
}

// NewCodeDescriptorMock returns a mock for CodeDescriptor
func NewCodeDescriptorMock(t minimock.Tester) *CodeDescriptorMock {
	m := &CodeDescriptorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CodeMock = mCodeDescriptorMockCode{mock: m}

	m.MachineTypeMock = mCodeDescriptorMockMachineType{mock: m}

	m.RefMock = mCodeDescriptorMockRef{mock: m}

	return m
}

type mCodeDescriptorMockCode struct {
	mock               *CodeDescriptorMock
	defaultExpectation *CodeDescriptorMockCodeExpectation
	expectations       []*CodeDescriptorMockCodeExpectation
}

// CodeDescriptorMockCodeExpectation specifies expectation struct of the CodeDescriptor.Code
type CodeDescriptorMockCodeExpectation struct {
	mock *CodeDescriptorMock

	results *CodeDescriptorMockCodeResults
	Counter uint64
}

// CodeDescriptorMockCodeResults contains results of the CodeDescriptor.Code
type CodeDescriptorMockCodeResults struct {
	ba1 []byte
	err error
}

// Expect sets up expected params for CodeDescriptor.Code
func (mmCode *mCodeDescriptorMockCode) Expect() *mCodeDescriptorMockCode {
	if mmCode.mock.funcCode != nil {
		mmCode.mock.t.Fatalf("CodeDescriptorMock.Code mock is already set by Set")
	}

	if mmCode.defaultExpectation == nil {
		mmCode.defaultExpectation = &CodeDescriptorMockCodeExpectation{}
	}

	return mmCode
}

// Inspect accepts an inspector function that has same arguments as the CodeDescriptor.Code
func (mmCode *mCodeDescriptorMockCode) Inspect(f func()) *mCodeDescriptorMockCode {
	if mmCode.mock.inspectFuncCode != nil {
		mmCode.mock.t.Fatalf("Inspect function is already set for CodeDescriptorMock.Code")
	}

	mmCode.mock.inspectFuncCode = f

	return mmCode
}

// Return sets up results that will be returned by CodeDescriptor.Code
func (mmCode *mCodeDescriptorMockCode) Return(ba1 []byte, err error) *CodeDescriptorMock {
	if mmCode.mock.funcCode != nil {
		mmCode.mock.t.Fatalf("CodeDescriptorMock.Code mock is already set by Set")
	}

	if mmCode.defaultExpectation == nil {
		mmCode.defaultExpectation = &CodeDescriptorMockCodeExpectation{mock: mmCode.mock}
	}
	mmCode.defaultExpectation.results = &CodeDescriptorMockCodeResults{ba1, err}
	return mmCode.mock
}

//Set uses given function f to mock the CodeDescriptor.Code method
func (mmCode *mCodeDescriptorMockCode) Set(f func() (ba1 []byte, err error)) *CodeDescriptorMock {
	if mmCode.defaultExpectation != nil {
		mmCode.mock.t.Fatalf("Default expectation is already set for the CodeDescriptor.Code method")
	}

	if len(mmCode.expectations) > 0 {
		mmCode.mock.t.Fatalf("Some expectations are already set for the CodeDescriptor.Code method")
	}

	mmCode.mock.funcCode = f
	return mmCode.mock
}

// Code implements CodeDescriptor
func (mmCode *CodeDescriptorMock) Code() (ba1 []byte, err error) {
	mm_atomic.AddUint64(&mmCode.beforeCodeCounter, 1)
	defer mm_atomic.AddUint64(&mmCode.afterCodeCounter, 1)

	if mmCode.inspectFuncCode != nil {
		mmCode.inspectFuncCode()
	}

	if mmCode.CodeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCode.CodeMock.defaultExpectation.Counter, 1)

		mm_results := mmCode.CodeMock.defaultExpectation.results
		if mm_results == nil {
			mmCode.t.Fatal("No results are set for the CodeDescriptorMock.Code")
		}
		return (*mm_results).ba1, (*mm_results).err
	}
	if mmCode.funcCode != nil {
		return mmCode.funcCode()
	}
	mmCode.t.Fatalf("Unexpected call to CodeDescriptorMock.Code.")
	return
}

// CodeAfterCounter returns a count of finished CodeDescriptorMock.Code invocations
func (mmCode *CodeDescriptorMock) CodeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCode.afterCodeCounter)
}

// CodeBeforeCounter returns a count of CodeDescriptorMock.Code invocations
func (mmCode *CodeDescriptorMock) CodeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCode.beforeCodeCounter)
}

// MinimockCodeDone returns true if the count of the Code invocations corresponds
// the number of defined expectations
func (m *CodeDescriptorMock) MinimockCodeDone() bool {
	for _, e := range m.CodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCode != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		return false
	}
	return true
}

// MinimockCodeInspect logs each unmet expectation
func (m *CodeDescriptorMock) MinimockCodeInspect() {
	for _, e := range m.CodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CodeDescriptorMock.Code")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		m.t.Error("Expected call to CodeDescriptorMock.Code")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCode != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		m.t.Error("Expected call to CodeDescriptorMock.Code")
	}
}

type mCodeDescriptorMockMachineType struct {
	mock               *CodeDescriptorMock
	defaultExpectation *CodeDescriptorMockMachineTypeExpectation
	expectations       []*CodeDescriptorMockMachineTypeExpectation
}

// CodeDescriptorMockMachineTypeExpectation specifies expectation struct of the CodeDescriptor.MachineType
type CodeDescriptorMockMachineTypeExpectation struct {
	mock *CodeDescriptorMock

	results *CodeDescriptorMockMachineTypeResults
	Counter uint64
}

// CodeDescriptorMockMachineTypeResults contains results of the CodeDescriptor.MachineType
type CodeDescriptorMockMachineTypeResults struct {
	m1 insolar.MachineType
}

// Expect sets up expected params for CodeDescriptor.MachineType
func (mmMachineType *mCodeDescriptorMockMachineType) Expect() *mCodeDescriptorMockMachineType {
	if mmMachineType.mock.funcMachineType != nil {
		mmMachineType.mock.t.Fatalf("CodeDescriptorMock.MachineType mock is already set by Set")
	}

	if mmMachineType.defaultExpectation == nil {
		mmMachineType.defaultExpectation = &CodeDescriptorMockMachineTypeExpectation{}
	}

	return mmMachineType
}

// Inspect accepts an inspector function that has same arguments as the CodeDescriptor.MachineType
func (mmMachineType *mCodeDescriptorMockMachineType) Inspect(f func()) *mCodeDescriptorMockMachineType {
	if mmMachineType.mock.inspectFuncMachineType != nil {
		mmMachineType.mock.t.Fatalf("Inspect function is already set for CodeDescriptorMock.MachineType")
	}

	mmMachineType.mock.inspectFuncMachineType = f

	return mmMachineType
}

// Return sets up results that will be returned by CodeDescriptor.MachineType
func (mmMachineType *mCodeDescriptorMockMachineType) Return(m1 insolar.MachineType) *CodeDescriptorMock {
	if mmMachineType.mock.funcMachineType != nil {
		mmMachineType.mock.t.Fatalf("CodeDescriptorMock.MachineType mock is already set by Set")
	}

	if mmMachineType.defaultExpectation == nil {
		mmMachineType.defaultExpectation = &CodeDescriptorMockMachineTypeExpectation{mock: mmMachineType.mock}
	}
	mmMachineType.defaultExpectation.results = &CodeDescriptorMockMachineTypeResults{m1}
	return mmMachineType.mock
}

//Set uses given function f to mock the CodeDescriptor.MachineType method
func (mmMachineType *mCodeDescriptorMockMachineType) Set(f func() (m1 insolar.MachineType)) *CodeDescriptorMock {
	if mmMachineType.defaultExpectation != nil {
		mmMachineType.mock.t.Fatalf("Default expectation is already set for the CodeDescriptor.MachineType method")
	}

	if len(mmMachineType.expectations) > 0 {
		mmMachineType.mock.t.Fatalf("Some expectations are already set for the CodeDescriptor.MachineType method")
	}

	mmMachineType.mock.funcMachineType = f
	return mmMachineType.mock
}

// MachineType implements CodeDescriptor
func (mmMachineType *CodeDescriptorMock) MachineType() (m1 insolar.MachineType) {
	mm_atomic.AddUint64(&mmMachineType.beforeMachineTypeCounter, 1)
	defer mm_atomic.AddUint64(&mmMachineType.afterMachineTypeCounter, 1)

	if mmMachineType.inspectFuncMachineType != nil {
		mmMachineType.inspectFuncMachineType()
	}

	if mmMachineType.MachineTypeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmMachineType.MachineTypeMock.defaultExpectation.Counter, 1)

		mm_results := mmMachineType.MachineTypeMock.defaultExpectation.results
		if mm_results == nil {
			mmMachineType.t.Fatal("No results are set for the CodeDescriptorMock.MachineType")
		}
		return (*mm_results).m1
	}
	if mmMachineType.funcMachineType != nil {
		return mmMachineType.funcMachineType()
	}
	mmMachineType.t.Fatalf("Unexpected call to CodeDescriptorMock.MachineType.")
	return
}

// MachineTypeAfterCounter returns a count of finished CodeDescriptorMock.MachineType invocations
func (mmMachineType *CodeDescriptorMock) MachineTypeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMachineType.afterMachineTypeCounter)
}

// MachineTypeBeforeCounter returns a count of CodeDescriptorMock.MachineType invocations
func (mmMachineType *CodeDescriptorMock) MachineTypeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMachineType.beforeMachineTypeCounter)
}

// MinimockMachineTypeDone returns true if the count of the MachineType invocations corresponds
// the number of defined expectations
func (m *CodeDescriptorMock) MinimockMachineTypeDone() bool {
	for _, e := range m.MachineTypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MachineTypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMachineTypeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMachineType != nil && mm_atomic.LoadUint64(&m.afterMachineTypeCounter) < 1 {
		return false
	}
	return true
}

// MinimockMachineTypeInspect logs each unmet expectation
func (m *CodeDescriptorMock) MinimockMachineTypeInspect() {
	for _, e := range m.MachineTypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CodeDescriptorMock.MachineType")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MachineTypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMachineTypeCounter) < 1 {
		m.t.Error("Expected call to CodeDescriptorMock.MachineType")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMachineType != nil && mm_atomic.LoadUint64(&m.afterMachineTypeCounter) < 1 {
		m.t.Error("Expected call to CodeDescriptorMock.MachineType")
	}
}

type mCodeDescriptorMockRef struct {
	mock               *CodeDescriptorMock
	defaultExpectation *CodeDescriptorMockRefExpectation
	expectations       []*CodeDescriptorMockRefExpectation
}

// CodeDescriptorMockRefExpectation specifies expectation struct of the CodeDescriptor.Ref
type CodeDescriptorMockRefExpectation struct {
	mock *CodeDescriptorMock

	results *CodeDescriptorMockRefResults
	Counter uint64
}

// CodeDescriptorMockRefResults contains results of the CodeDescriptor.Ref
type CodeDescriptorMockRefResults struct {
	rp1 *insolar.Reference
}

// Expect sets up expected params for CodeDescriptor.Ref
func (mmRef *mCodeDescriptorMockRef) Expect() *mCodeDescriptorMockRef {
	if mmRef.mock.funcRef != nil {
		mmRef.mock.t.Fatalf("CodeDescriptorMock.Ref mock is already set by Set")
	}

	if mmRef.defaultExpectation == nil {
		mmRef.defaultExpectation = &CodeDescriptorMockRefExpectation{}
	}

	return mmRef
}

// Inspect accepts an inspector function that has same arguments as the CodeDescriptor.Ref
func (mmRef *mCodeDescriptorMockRef) Inspect(f func()) *mCodeDescriptorMockRef {
	if mmRef.mock.inspectFuncRef != nil {
		mmRef.mock.t.Fatalf("Inspect function is already set for CodeDescriptorMock.Ref")
	}

	mmRef.mock.inspectFuncRef = f

	return mmRef
}

// Return sets up results that will be returned by CodeDescriptor.Ref
func (mmRef *mCodeDescriptorMockRef) Return(rp1 *insolar.Reference) *CodeDescriptorMock {
	if mmRef.mock.funcRef != nil {
		mmRef.mock.t.Fatalf("CodeDescriptorMock.Ref mock is already set by Set")
	}

	if mmRef.defaultExpectation == nil {
		mmRef.defaultExpectation = &CodeDescriptorMockRefExpectation{mock: mmRef.mock}
	}
	mmRef.defaultExpectation.results = &CodeDescriptorMockRefResults{rp1}
	return mmRef.mock
}

//Set uses given function f to mock the CodeDescriptor.Ref method
func (mmRef *mCodeDescriptorMockRef) Set(f func() (rp1 *insolar.Reference)) *CodeDescriptorMock {
	if mmRef.defaultExpectation != nil {
		mmRef.mock.t.Fatalf("Default expectation is already set for the CodeDescriptor.Ref method")
	}

	if len(mmRef.expectations) > 0 {
		mmRef.mock.t.Fatalf("Some expectations are already set for the CodeDescriptor.Ref method")
	}

	mmRef.mock.funcRef = f
	return mmRef.mock
}

// Ref implements CodeDescriptor
func (mmRef *CodeDescriptorMock) Ref() (rp1 *insolar.Reference) {
	mm_atomic.AddUint64(&mmRef.beforeRefCounter, 1)
	defer mm_atomic.AddUint64(&mmRef.afterRefCounter, 1)

	if mmRef.inspectFuncRef != nil {
		mmRef.inspectFuncRef()
	}

	if mmRef.RefMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRef.RefMock.defaultExpectation.Counter, 1)

		mm_results := mmRef.RefMock.defaultExpectation.results
		if mm_results == nil {
			mmRef.t.Fatal("No results are set for the CodeDescriptorMock.Ref")
		}
		return (*mm_results).rp1
	}
	if mmRef.funcRef != nil {
		return mmRef.funcRef()
	}
	mmRef.t.Fatalf("Unexpected call to CodeDescriptorMock.Ref.")
	return
}

// RefAfterCounter returns a count of finished CodeDescriptorMock.Ref invocations
func (mmRef *CodeDescriptorMock) RefAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRef.afterRefCounter)
}

// RefBeforeCounter returns a count of CodeDescriptorMock.Ref invocations
func (mmRef *CodeDescriptorMock) RefBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRef.beforeRefCounter)
}

// MinimockRefDone returns true if the count of the Ref invocations corresponds
// the number of defined expectations
func (m *CodeDescriptorMock) MinimockRefDone() bool {
	for _, e := range m.RefMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RefMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRefCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRef != nil && mm_atomic.LoadUint64(&m.afterRefCounter) < 1 {
		return false
	}
	return true
}

// MinimockRefInspect logs each unmet expectation
func (m *CodeDescriptorMock) MinimockRefInspect() {
	for _, e := range m.RefMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CodeDescriptorMock.Ref")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RefMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRefCounter) < 1 {
		m.t.Error("Expected call to CodeDescriptorMock.Ref")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRef != nil && mm_atomic.LoadUint64(&m.afterRefCounter) < 1 {
		m.t.Error("Expected call to CodeDescriptorMock.Ref")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CodeDescriptorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCodeInspect()

		m.MinimockMachineTypeInspect()

		m.MinimockRefInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CodeDescriptorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CodeDescriptorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCodeDone() &&
		m.MinimockMachineTypeDone() &&
		m.MinimockRefDone()
}
