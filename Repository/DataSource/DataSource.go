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


// DataSource project DataSource.go
package DataSource

import (
    . "github.com/ivan-kostko/GoLibs/CustomErrors"
)

// Represents the instruction to DataSource
// For SQL DataSource it would be a SQL-script; for http service DataSource it would be an access link; and so on
type Instruction string

// Represents a result of the instruction
type Result interface{}

// Represents standard instruction execution function
type ExecuteInstruction func(i Instruction) (Result, *Error)

// Represents generic data source or persistence
type DataSource struct{
    ExecuteInstruction
}

// Represents generic DataSource factory
func GetNewDataSource(ei ExecuteInstruction) *DataSource {
    return &DataSource{ExecuteInstruction: ei}
}
