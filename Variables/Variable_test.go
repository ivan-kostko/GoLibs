package Variables

import (
	"reflect"
	"testing"
)

type MockRWLock struct {
	logChan  chan string
	doneChan chan struct{}
	callLog  []string
}

func (this *MockRWLock) addToLog(s string) {
	this.logChan <- s
}

func (this *MockRWLock) Lock() {
	this.addToLog("Lock")
}

func (this *MockRWLock) RLock() {
	this.addToLog("RLock")
}

func (this *MockRWLock) Unlock() {
	this.addToLog("Unlock")
}

func (this *MockRWLock) RUnlock() {
	this.addToLog("RUnlock")
}

func NewMockRWLock() *MockRWLock {

	m := &MockRWLock{
		logChan:  make(chan string),
		callLog:  make([]string, 0),
		doneChan: make(chan struct{}),
	}

	go func(m *MockRWLock) {
		defer close(m.doneChan)

		for s := range m.logChan {
			m.callLog = append(m.callLog, s)
		}
	}(m)

	return m
}

func (this *MockRWLock) ShutDown() []string {
	close(this.logChan)
	for _ = range this.doneChan {

	}
	return this.callLog
}

func Test_NewVariable(t *testing.T) {

	var v *Variable = New()

	if v == nil {
		t.Error("New() returned nil")
	}

	if v.lock == nil {
		t.Error("New() returned *Variable with nil lock")
		t.FailNow()
	}
}

func Test_SetInterfaceValue(t *testing.T) {

	testCases := []struct {
		TestAlias     string
		SetValue      interface{}
		ExpectedValue interface{}
	}{
		{
			TestAlias:     "Set Value to nil simple call",
			SetValue:      nil,
			ExpectedValue: nil,
		},
		{
			TestAlias:     "Set Value to 1 simple call",
			SetValue:      1,
			ExpectedValue: 1,
		},
		{
			TestAlias:     "Set Value to \"xsd\" simple call",
			SetValue:      "xsd",
			ExpectedValue: "xsd",
		},
		{
			TestAlias:     "Set Value to map simple call",
			SetValue:      map[string]interface{}{"dsx": "xsd"},
			ExpectedValue: map[string]interface{}{"dsx": "xsd"},
		},
	}

	for _, testCase := range testCases {

		testAlias := testCase.TestAlias
		setValue := testCase.SetValue
		expectedValue := testCase.ExpectedValue

		testFn := func(t *testing.T) {

			v := New()

			v.SetValue(setValue)

			actualValue := v.InterfaceValue()
			if !reflect.DeepEqual(actualValue, expectedValue) {
				t.Errorf("\r\n %s :: SetValue(%#v) set value to \r\n %+v \r\n while expected \r\n %+v \r\n", testAlias, setValue, actualValue, expectedValue)

			}
		}
		t.Run(testAlias, testFn)
	}

}

func Test_RWLockUsage(t *testing.T) {

	testCases := []struct {
		TestAlias       string
		VOperations     func(v *Variable)
		ExpectedLockLog []string
	}{
		{
			TestAlias:       "Set only. nil",
			VOperations:     func(v *Variable) { v.SetValue(nil) },
			ExpectedLockLog: []string{},
		},
		{
			TestAlias:       "Set only. \"xsd\"",
			VOperations:     func(v *Variable) { v.SetValue("xsd") },
			ExpectedLockLog: []string{"Lock", "Unlock"},
		},
		{
			TestAlias:       "Get only. nil",
			VOperations:     func(v *Variable) { v.InterfaceValue() },
			ExpectedLockLog: []string{"RLock", "RUnlock"},
		},
		{
			TestAlias:       "Set nil, Get, Set \"xsd\"",
			VOperations:     func(v *Variable) { v.SetValue(nil); v.InterfaceValue(); v.SetValue("xsd") },
			ExpectedLockLog: []string{"RLock", "RUnlock", "Lock", "Unlock"},
		},
	}

	for _, testCase := range testCases {

		testAlias := testCase.TestAlias
		vOperations := testCase.VOperations
		expectedLockLog := testCase.ExpectedLockLog

		testFn := func(t *testing.T) {

			lock := NewMockRWLock()
			v := New()
			v.lock = lock

			vOperations(v)

			actualLockLog := lock.ShutDown()
			if !reflect.DeepEqual(actualLockLog, expectedLockLog) {
				t.Errorf("\r\n %s :: LockLog after VOperations invokation is like the following \r\n %+v \r\n while expected \r\n %+v \r\n", testAlias, actualLockLog, expectedLockLog)

			}
		}
		t.Run(testAlias, testFn)
	}
}

func Test_OnLockalChangeHookedTriggersInvokation(t *testing.T) {

	actualInvokationLog := make([]string, 0)

	testCases := []struct {
		TestAlias             string
		VOperations           func(v *Variable)
		ExpectedInvokationLog []string
	}{
		{
			TestAlias: "Set only. nil",
			VOperations: func(v *Variable) {
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 1") })
				v.SetValue(nil)
			},
			ExpectedInvokationLog: []string{},
		},
		{
			TestAlias: "Set only. \"xsd\"",
			VOperations: func(v *Variable) {
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 1") })
				v.SetValue("xsd")
			},
			ExpectedInvokationLog: []string{"Trigger 1"},
		},
		{
			TestAlias:             "Get only. nil",
			VOperations:           func(v *Variable) { v.InterfaceValue() },
			ExpectedInvokationLog: []string{},
		},
		{
			TestAlias: "Set nil, Get, Set \"xsd\"",
			VOperations: func(v *Variable) {
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 1") })
				v.SetValue(nil)
				v.InterfaceValue()
				v.SetValue("xsd")
			},
			ExpectedInvokationLog: []string{"Trigger 1"},
		},
		{
			TestAlias: "Set \"xsd\" to 1 with 3 triggers",
			VOperations: func(v *Variable) {
				v.SetValue("xsd")
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 1") })
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 2") })
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 3") })
				v.SetValue(1)

			},
			ExpectedInvokationLog: []string{"Trigger 1", "Trigger 2", "Trigger 3"},
		},
		{
			TestAlias: "Set \"xsd\" to 1 and then to []string{\"X\", \"s\", \"d\"} with 3 triggers",
			VOperations: func(v *Variable) {
				v.SetValue("xsd")
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 1") })
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 2") })
				v.HookupAfterLocalChangeTrigger(func(v *Variable) { actualInvokationLog = append(actualInvokationLog, "Trigger 3") })
				v.SetValue(1)
				v.SetValue([]string{"X", "s", "d"})

			},
			ExpectedInvokationLog: []string{"Trigger 1", "Trigger 2", "Trigger 3", "Trigger 1", "Trigger 2", "Trigger 3"},
		},
	}

	for _, testCase := range testCases {

		testAlias := testCase.TestAlias
		vOperations := testCase.VOperations
		expectedInvokationLog := testCase.ExpectedInvokationLog

		testFn := func(t *testing.T) {

			actualInvokationLog = []string{}

			v := New()
			vOperations(v)

			if !reflect.DeepEqual(actualInvokationLog, expectedInvokationLog) {
				t.Errorf("\r\n %s :: InvokationLog after VOperations invokation is like the following \r\n %+v \r\n while expected \r\n %+v \r\n", testAlias, actualInvokationLog, expectedInvokationLog)

			}
		}
		t.Run(testAlias, testFn)
	}

}

func Test_InterfaceValue(t *testing.T) {

	testCases := []struct {
		TestAlias     string
		Value         interface{}
		ExpectedValue interface{}
	}{
		{
			TestAlias:     "Set Value to nil simple call",
			Value:         nil,
			ExpectedValue: nil,
		},
		{
			TestAlias:     "Set Value to 1 simple call",
			Value:         1,
			ExpectedValue: 1,
		},
		{
			TestAlias:     "Set Value to \"xsd\" simple call",
			Value:         "xsd",
			ExpectedValue: "xsd",
		},
		{
			TestAlias:     "Set Value to map simple call",
			Value:         map[string]interface{}{"dsx": "xsd"},
			ExpectedValue: map[string]interface{}{"dsx": "xsd"},
		},
	}

	for _, testCase := range testCases {

		testAlias := testCase.TestAlias
		value := testCase.Value
		expectedValue := testCase.ExpectedValue

		testFn := func(t *testing.T) {

			v := New()
			v.SetValue(value)

			actualValue := v.InterfaceValue()

			if !reflect.DeepEqual(actualValue, expectedValue) {
				t.Errorf("\r\n %s :: v.InterfaceValue() returned \r\n %+v \r\n while expected \r\n %+v \r\n", testAlias, actualValue, expectedValue)

			}

		}
		t.Run(testAlias, testFn)
	}

}

func Test_HookupAfterLocalChangeTriggerGotV(t *testing.T) {

	v := New()

	var trigger1GotV, trigger2GotV, trigger3GotV *Variable
	v.HookupAfterLocalChangeTrigger(func(v *Variable) { trigger1GotV = v })
	v.HookupAfterLocalChangeTrigger(func(v *Variable) { trigger2GotV = v })
	v.HookupAfterLocalChangeTrigger(func(v *Variable) { trigger3GotV = v })

	v.SetValue(1)

	if trigger1GotV != v {
		t.Errorf("Trigger1 got %#v while expected %#v", trigger1GotV, v)
	}
	if trigger2GotV != v {
		t.Errorf("Trigger2 got %#v while expected %#v", trigger2GotV, v)
	}
	if trigger3GotV != v {
		t.Errorf("Trigger3 got %#v while expected %#v", trigger3GotV, v)
	}

}
