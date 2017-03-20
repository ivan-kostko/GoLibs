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

import "fmt"

// Represents implementation limitations constants
const (
	// Represents the limit of considered nested function calls
	MAXNESTEDLEVEL = 128
)

// Represents goroutnes function call stack
type callStack []struct {
	FullFileName string
	Line         int
	FunctionName string
}

// Represents callStack factory and its DEFAULT implementation, which returns nil (recommended for production).
// Generates up to MAXNESTEDLEVEL entries starting from skip element calling this function.
// So, in case skip is 0, it will generate call stack starting from newCallStack caller.
// Nevertheless, it is overriden in fullCallStack and justCaller if debugging is needed.
var newCallStack = func(skip int) callStack { return nil }

func (this callStack) toString() (s string) {
	for i := 0; i < len(this); i++ {
		call := this[i]
		s += fmt.Sprintln(call.FullFileName, call.Line, call.FunctionName)
	}

	return
}
