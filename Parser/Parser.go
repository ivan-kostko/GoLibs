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

const INIT_PARSERSCAPACITY = 10

// Represents the list of registered parsers
var parsers = tsMap.New(INIT_PARSERSCAPACITY)

// Defines functionality of parser as combination of two functions: Serialize + Deserialize
type Parser struct {
	Serializer
	Deserializer
}

// Registers parser for a format
// In case there is already registered Parser with same format it returns ERR_ALREADYREGISTERED skipping further registration steps. So, initial registration stays w/o changes.
// For.Ex. having registered SomeFormat + SomeParser, registration of SomeFormat + NewParser returns ERR_ALREADYREGISTERED error, keeping initial registration (SomeFormat + SomeParser) w/o changes.
func Register(f FormatParser, p *Parser) *Error {
	if _, ok := parsers.Get(f.String()); ok {
		return NewError(InvalidOperation, ERR_ALREADYREGISTERED)
	}
	parsers.Set(f.String(), p)
	return nil
}

// Gets parser by format name as string
// In case of error returns nil and InvalidOperation error with one of predefined messages:
// ERR_WONTGETPARSER
// ERR_WRONGREGTYPE
func getParserByFormatName(f string) (parser *Parser, err *Error) {
	p, ok := parsers.Get(f)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WONTGETPARSER)
	}
	parser, ok = p.(*Parser)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WRONGREGTYPE)
	}
	return parser, nil
}

// Gets parser by format
// In case of error returns empty parser and InvalidOperation error with one of predefined messages:
// ERR_WONTGETPARSER
// ERR_WRONGREGTYPE
func GetParserByFormat(f FormatParser) (parser *Parser, err *Error) {
	return getParserByFormatName(f.String())
}
