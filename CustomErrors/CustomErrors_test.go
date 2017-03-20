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
	"fmt"
	"testing"
)

func Example() {
	// Define a function which will return *Error
	fn := func(i interface{}) *Error {
		// It will print if i is string
		// and return Error with type InvalidArgument - if not.
		switch t := i.(type) {
		case string:

			// Print out
			fmt.Println("Yup, it is a string")
			return nil
			break
		default:

			// Here we prepare complete and detailed error message
			msg := fmt.Sprintf("Input parameter is of type %T while acceptable is only string", t)

			return NewError(InvalidArgument, msg)
		}
		return nil
	}

	i := 456
	err := fn(i)
	if err != nil {
		if err.Type == InvalidArgument {

			// Print Error message
			fmt.Println(err.Message)

			// Now we could explicitly convert i into string and pass it again
			_ = fn(fmt.Sprint(i))
		} else {
			// ... do something
		}
	}
	// Output: Input parameter is of type int while acceptable is only string
	// Yup, it is a string
}

func TestErrorImplements_errorInterface(t *testing.T) {
	var _ error = &Error{}
}

func TestErrorTypeStringer(t *testing.T) {

	testCases := []struct {
		TestAlias      string
		ErrorType      ErrorType
		ExpectedString string
	}{
		{
			TestAlias:      "BasicError",
			ErrorType:      BasicError,
			ExpectedString: "BasicError",
		},
		{
			TestAlias:      "InvalidOperation",
			ErrorType:      InvalidOperation,
			ExpectedString: "InvalidOperation",
		},
		{
			TestAlias:      "InvalidArgument",
			ErrorType:      InvalidArgument,
			ExpectedString: "InvalidArgument",
		},
		{
			TestAlias:      "AccessViolation",
			ErrorType:      AccessViolation,
			ExpectedString: "AccessViolation",
		},
		{
			TestAlias:      "Nonsupported",
			ErrorType:      Nonsupported,
			ExpectedString: "Nonsupported",
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		errorType := testCase.ErrorType
		expectedString := testCase.ExpectedString

		testFn := func(t *testing.T) {

			actualString := fmt.Sprintf("%s", errorType)
			if actualString != expectedString {
				t.Errorf("%s :: Error.Error() returned\r\n %v \r\n while expected \r\n %v", testAlias, actualString, expectedString)
			}

		}
		t.Run(testAlias, testFn)
	}

}

func TestErrorErrorMethod(t *testing.T) {

	testCases := []struct {
		TestAlias      string
		ActualError    Error
		ExpectedString string
	}{
		{
			TestAlias:      `CustomErrors.Error{Type:InvalidArgument, Message:Testing}`,
			ActualError:    Error{Type: InvalidArgument, Message: "Testing"},
			ExpectedString: `CustomErrors.Error{Type:InvalidArgument, Message:Testing}`,
		},
		{
			TestAlias:      `CustomErrors.Error{Type:InvalidOperation, Message:Testing}`,
			ActualError:    Error{Type: InvalidOperation, Message: "Testing"},
			ExpectedString: `CustomErrors.Error{Type:InvalidOperation, Message:Testing}`,
		},
		{
			TestAlias:      `CustomErrors.Error{Type:AccessViolation, Message:Testing}`,
			ActualError:    Error{Type: AccessViolation, Message: "Testing"},
			ExpectedString: `CustomErrors.Error{Type:AccessViolation, Message:Testing}`,
		},
	}

	for _, testCase := range testCases {
		testAlias := testCase.TestAlias
		actualError := testCase.ActualError
		expectedString := testCase.ExpectedString

		testFn := func(t *testing.T) {

			actualString := actualError.Error()
			if actualString != expectedString {
				t.Errorf("%s :: Error.Error() returned\r\n %v \r\n while expected \r\n %v", testAlias, actualString, expectedString)
			}

		}
		t.Run(testAlias, testFn)
	}
}
