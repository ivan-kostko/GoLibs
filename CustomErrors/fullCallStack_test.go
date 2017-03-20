//+build fullCallStack

//   Copyright (c) 2016 Ivan A Kostko (github.com/ivan-kostko)

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//       http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Errors project Errors.go
// The package contains general error exteded functionality
package CustomErrors

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_fullCallStack(t *testing.T) {

	var actualStackRepresentation callStack

	pwd, err := os.Getwd()
	if err != nil {
		t.Log(err.Error())
	}

	testCases := []struct {
		TestAlias                   string
		CallFunctionsStack          func()
		ExpectedStackRepresentation callStack
	}{
		{
			TestAlias:          "Simple not nested call with skip 0",
			CallFunctionsStack: func() { actualStackRepresentation = newFullCallStack(0) },
			ExpectedStackRepresentation: []struct {
				FullFileName string
				Line         int
				FunctionName string
			}{
				{
					FullFileName: strings.TrimRight(pwd, "/") + "/fullCallStack_test.go",
					Line:         44,
					FunctionName: "GoLibs/CustomErrors.Test_fullCallStack.func1",
				},
				{
					FullFileName: strings.TrimRight(pwd, "/") + "/fullCallStack_test.go",
					Line:         79,
					FunctionName: "GoLibs/CustomErrors.Test_fullCallStack",
				},
				{
					FullFileName: strings.TrimRight(os.Getenv("GOROOT"), "/") + "/src/testing/testing.go",
					Line:         610,
					FunctionName: "testing.tRunner",
				},
				{
					FullFileName: strings.TrimRight(os.Getenv("GOROOT"), "/") + "/src/runtime/asm_amd64.s",
					Line:         2086,
					FunctionName: "runtime.goexit",
				},
			},
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		callFunctionsStack := testCase.CallFunctionsStack
		expectedStackRepresentation := testCase.ExpectedStackRepresentation

		callFunctionsStack()

		testFn := func(t *testing.T) {

			if !reflect.DeepEqual(actualStackRepresentation, expectedStackRepresentation) {
				t.Errorf("%s :: returned \r\n %s \r\n while expected \r\n %s", testAlias, actualStackRepresentation, expectedStackRepresentation)
			}
		}

		t.Run(testAlias, testFn)
	}
}

func Test_NewErrorWithCallStack(t *testing.T) {

	var actualError *Error
	pwd, err := os.Getwd()
	if err != nil {
		t.Log(err.Error())
	}

	testCases := []struct {
		TestAlias          string
		CallFunctionsStack func()
		ExpectedError      *Error
	}{
		{
			TestAlias:          "Not nested BasicError",
			CallFunctionsStack: func() { actualError = NewError(BasicError, "Basic Error message") },
			ExpectedError: &Error{
				Type:    BasicError,
				Message: "Basic Error message",
				callStack: []struct {
					FullFileName string
					Line         int
					FunctionName string
				}{
					{
						FullFileName: strings.TrimRight(pwd, "/") + "/fullCallStack_test.go",
						Line:         107,
						FunctionName: "GoLibs/CustomErrors.Test_NewErrorWithCallStack.func1",
					},
					{
						FullFileName: strings.TrimRight(pwd, "/") + "/fullCallStack_test.go",
						Line:         146,
						FunctionName: "GoLibs/CustomErrors.Test_NewErrorWithCallStack",
					},
					{
						FullFileName: strings.TrimRight(os.Getenv("GOROOT"), "/") + "/src/testing/testing.go",
						Line:         610,
						FunctionName: "testing.tRunner",
					},
					{
						FullFileName: strings.TrimRight(os.Getenv("GOROOT"), "/") + "/src/runtime/asm_amd64.s",
						Line:         2086,
						FunctionName: "runtime.goexit",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		callFunctionsStack := testCase.CallFunctionsStack
		expectedError := testCase.ExpectedError

		callFunctionsStack()

		testFn := func(t *testing.T) {

			if !reflect.DeepEqual(actualError, expectedError) {
				t.Errorf("%s :: returned \r\n %#v \r\n while expected \r\n %#v", testAlias, actualError, expectedError)
			}
		}

		t.Run(testAlias, testFn)
	}
}

func TestErrorMethodWithCallStack(t *testing.T) {

	testCases := []struct {
		TestAlias           string
		ActualError         *Error
		ExpectedErrorString string
	}{
		{
			TestAlias: "BasicError with nil call stack",
			ActualError: &Error{
				Type:    BasicError,
				Message: "Basic Error message",
			},
			ExpectedErrorString: `CustomErrors.Error{Type:BasicError, Message:Basic Error message}`,
		},
		{
			TestAlias: "BasicError with nil call stack",
			ActualError: &Error{
				Type:    BasicError,
				Message: "Basic Error message",
				callStack: []struct {
					FullFileName string
					Line         int
					FunctionName string
				}{
					{
						FullFileName: "fileOne.go",
						Line:         107,
						FunctionName: "TestFunctionOne",
					},
					{
						FullFileName: "fileTwo.go",
						Line:         146,
						FunctionName: "TestFunctionTwo",
					},
				},
			},
			ExpectedErrorString: "CustomErrors.Error{Type:BasicError, Message:Basic Error message\r\nCall Stack: \r\n fileOne.go 107 TestFunctionOne\nfileTwo.go 146 TestFunctionTwo\n}",
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		actualError := testCase.ActualError
		expectedErrorString := testCase.ExpectedErrorString

		testFn := func(t *testing.T) {

			actualErrorString := actualError.Error()

			if actualErrorString != expectedErrorString {
				t.Errorf("%s :: returned \r\n %#v \r\n while expected \r\n %#v", testAlias, actualErrorString, expectedErrorString)
			}
		}

		t.Run(testAlias, testFn)
	}
}
