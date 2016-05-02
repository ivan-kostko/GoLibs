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

package Default

import (
	"encoding/xml"

	. "github.com/ivan-kostko/GoLibs/CustomErrors"
	parsers "github.com/ivan-kostko/GoLibs/Parser"
)

var parser = parsers.Parser{Serializer, Deserializer}

func init() {
	parsers.Register(parsers.DefaultXML, &parser)
}

func Serializer(in interface{}) ([]byte, *Error) {
	b, err := xml.Marshal(in)
	if err != nil {
		return nil, NewError(InvalidOperation, err.Error())
	}
	return b, nil
}

func Deserializer(document []byte, dest interface{}) *Error {
	err := xml.Unmarshal(document, dest)
	if err != nil {
		return NewError(InvalidOperation, err.Error())
	}
	return nil
}
