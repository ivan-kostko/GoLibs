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

package CustomErrors

import (
	"fmt"
)

//go:generate stringer -type=ErrorType

// Represents enum of predefined error types
type ErrorType int

const (
	BasicError ErrorType = iota
	InvalidOperation
	InvalidArgument
	AccessViolation
	Nonsupported
)

// Represents custom error as tuple Type + Message.
type Error struct {
	Type    ErrorType
	Message string
}

// Implementation of standart error interface
func (e Error) Error() string {
	return fmt.Sprintf("%T{Type:%s, Message:%s}", e, e.Type, e.Message)
}

// Error factory
func NewError(typ ErrorType, msg string) *Error {
	return &Error{
		Type:    typ,
		Message: msg,
	}
}

// Error factory generating message in fmt.Sprintf manner
func NewErrorF(typ ErrorType, baseMsg string, args ...interface{}) *Error {
	msg := fmt.Sprintf(baseMsg, args...)
	return NewError(typ, msg)
}
