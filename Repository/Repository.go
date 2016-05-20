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

import(
    ds "./DataSource"
    . "github.com/ivan-kostko/GoLibs/CustomErrors"
)

const(
    ERR_FAILEDTOGENERATEINSTRUCTION = "Repository: Failed to generate instruction %s"
)

// Represents a data repository abstraction.
// It is supposed to be an interface to request and retreive entities between domains.
type Repository struct {
    dataSource        *ds.DataSource
    instructor        Instructor
    checkError        func(*Error) bool
}

// An alias for models in domain
type DomainModel interface{}

// Retreives all entities conforming FilteringCondition(s)
func (rep *Repository) GetAll(fc ...FilteringCondition) ([]DomainModel, *Error){
    instruction, err := rep.instructor.GenerateInstruction(fc...)
    if rep.checkError(err) {
        return nil, NewError(InvalidOperation, ERR_FAILEDTOGENERATEINSTRUCTION)
    }
    rep.dataSource.ExecuteInstruction(instruction)
    return nil,nil
}
