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

package CustomErrors

import "runtime"

func init() {
	newCallStack = newFullCallStack
}

func newFullCallStack(skip int) callStack {
	callers := make([]uintptr, MAXNESTEDLEVEL)

	// skip 1 for runtime.Callers(see its documentation) and 1 for current function
	x := runtime.Callers(skip+1+1, callers)
	frames := runtime.CallersFrames(callers[:x])
	if frames == nil {
		return nil
	}
	cs := make(callStack, 0, x)
	for {
		frame, more := frames.Next()
		cs = append(cs, struct {
			FullFileName string
			Line         int
			FunctionName string
		}{
			FullFileName: frame.File,
			Line:         frame.Line,
			FunctionName: frame.Function,
		})
		if !more {
			break
		}
	}

	return cs
}
