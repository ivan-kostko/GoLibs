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

func TestErrorErrorMethod(t *testing.T) {
	expected := `CustomErrors.Error{Type:AccessViolation, Message:Testing}`
	var err error = Error{Type: AccessViolation, Message: "Testing"}
	if err.Error() != expected {
		t.Errorf("err.Error() returned %v while expected %v", err.Error(), expected)
	}
}

func TestErrorTypeStringer(t *testing.T) {
	expected := "AccessViolation"
	if fmt.Sprintf("%s", AccessViolation) != expected {
		t.Errorf("Returned %s while expected %v", AccessViolation, expected)
	}
}
