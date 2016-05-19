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


package Repository

// Represents comparrison operator ENUM
type Operator string

const(
    EQ Operator = "EQ"  // Equal
    NE Operator = "NE"  // Not equal
    IN Operator = "IN"  // In range
    GE Operator = "GE"  // Greater or equal
    // TODO(me): Extend with other operators
)

// Represents ingle predicate as operator + values
type Predicate struct {
    Operator
    Values    []interface{}
}

// Represents composition of fiels as key and predicates
type Criteria map[string][]Predicate
