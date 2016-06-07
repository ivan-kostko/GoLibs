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

// Represents comparison operator ENUM
type Operator string

const(
    NOTIN Operator = "NOTIN"  // NOT in range
    IN Operator = "IN"  // In range
    GE Operator = "GE"  // Greater or equal
    GT Operator = "GT"  // Greater than
    LE Operator = "LE"  // Less or equal
    LT Operator = "LT"  // Less than
)

// Represents single predicate as operator + value(s)
type Predicate struct {
    Operator
    Values    []interface{}
}

// Represents filtering creteria, which is applied at DataSource instruction. For.Ex. for SQL DataSource it would be converted by Instructor into WHERE clause.
// It is implemented as composition of fields(key(s)) and predicates(value)
type FilteringCondition map[string][]Predicate
