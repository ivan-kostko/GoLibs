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

package Parser

import (
	. "github.com/ivan-kostko/GoLibs/CustomErrors"
	tsMap "github.com/ivan-kostko/GoLibs/ThreadSafe/Map"
)

// Predefined list of error messages
const (
	ERR_WONTGETPARSER     = "Parser: Won't get parser for provided format, cause it is not registered"
	ERR_WRONGREGTYPE      = "Parser: Won't get parser for provided format, cause it is of wrong type"
	ERR_ALREADYREGISTERED = "Parser: There is already registered parser for provided format. Wont register twice"
)

const INIT_CAPACITY = 10

// Represents the list of registered parsers
var parsers = tsMap.New(INIT_CAPACITY)

// Defines functionality of parser as combination of two functions: Serialize + Deserialize
type Parser struct {
	Serializer
	Deserializer
}

// Registers parser implementation by alias
// In case there is already registered Parser with same alias it returns ERR_ALREADYREGISTERED skipping further registration steps. So, initial registration stays w/o changes.
// For.Ex. having registered "SomeImplementation" + SomeParser, registration of "SomeImplementation" + NewParser returns ERR_ALREADYREGISTERED error, keeping initial registration (SomeFormat + SomeParser) w/o changes.
func Register(implementationAlias string, p *Parser) *Error {
	if _, ok := parsers.Get(implementationAlias); ok {
		return NewError(InvalidOperation, ERR_ALREADYREGISTERED)
	}
	parsers.Set(implementationAlias, p)
	return nil
}

// Gets parser by implementation by alias
// In case of error returns nil and InvalidOperation error with one of predefined messages:
// ERR_WONTGETPARSER
// ERR_WRONGREGTYPE
func GetParser(implementationAlias string) (parser *Parser, err *Error) {
	p, ok := parsers.Get(implementationAlias)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WONTGETPARSER)
	}
	parser, ok = p.(*Parser)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WRONGREGTYPE)
	}
	return parser, nil
}
