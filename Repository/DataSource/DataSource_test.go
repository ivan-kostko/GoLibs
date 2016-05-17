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

package DataSource

import (
    "testing"
    . "github.com/ivan-kostko/GoLibs/CustomErrors"

)

func TestGetNewDataSourceAssignedExecuteInstruction(t *testing.T){

    expected := true
    var actual bool = false

    ei := ExecuteInstruction(func(i Instruction) (Result, *Error){ actual = true; return nil,nil})
    ds := GetNewDataSource(ei)

    // Due to complexity of comparing functions (it is possible to compate only their pointers or compare to nil)
    // I'm just checking, that by the end of DataSource.ExecuteInstruction invokation the mock (ei) function is called, which is indicated by setting actual to true
    ds.ExecuteInstruction("")
    if actual != expected {
        t.Errorf("GetNewDataSource(ei) returned ExecuteInstruction as %#v while expected %#v", actual, expected)
    }
}
