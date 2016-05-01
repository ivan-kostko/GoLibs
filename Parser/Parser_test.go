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
	"testing"

	"reflect"

	. "github.com/ivan-kostko/GoLibs/CustomErrors"
	tsMap "github.com/ivan-kostko/GoLibs/ThreadSafe/Map"
)

func TestGetParser(t *testing.T) {

	type fakeParser struct{}

	defaultYAMLParser := &Parser{}

	testCases := []struct {
		RegisteredKey   string
		RegisteredValue interface{}
		GetFormat       Format
		ExpectedParser  *Parser
		ExpectedError   *Error
	}{
		{
			"DefaultJSON",
			interface{}(&fakeParser{}),
			DefaultXML,
			nil,
			NewError(InvalidOperation, ERR_WONTGETPARSER),
		},
		{
			"DefaultXML",
			interface{}(&fakeParser{}),
			DefaultXML,
			nil,
			NewError(InvalidOperation, ERR_WRONGREGTYPE),
		},
		{
			"DefaultYAML",
			defaultYAMLParser,
			DefaultYAML,
			defaultYAMLParser,
			nil,
		},
	}

	for _, testCase := range testCases {
		parsers.Set(testCase.RegisteredKey, testCase.RegisteredValue)
		actualParser, actualError := GetParserByFormat(testCase.GetFormat)
		if !(reflect.DeepEqual(actualParser, testCase.ExpectedParser) &&
			reflect.DeepEqual(actualError, testCase.ExpectedError)) {
			t.Errorf("GetParserByFormat(%v) returned Parser as %v and Error as %v \r\n\t\t\t while expected Parser as %v and Error as %v", testCase.GetFormat, actualParser, actualError, testCase.ExpectedParser, testCase.ExpectedError)
		}
		// Reset parsers
		parsers = tsMap.New(INIT_PARSERSCAPACITY)

	}
}

func TestRegisterParser(t *testing.T) {

	type fakeParser struct {
		Serializer
		Deserializer
	}

	defaultXMLParser := &Parser{}
	defaultJSONParser := &Parser{}
	defaultYAMLParser := &Parser{}

	testCases := []struct {
		RegisterFormat Format
		RegisterParser *Parser
		ExpectedParser *Parser
		ExpectedError  *Error
	}{
		{
			DefaultXML,
			defaultXMLParser,
			defaultXMLParser,
			nil,
		},
		{
			DefaultJSON,
			defaultJSONParser,
			defaultJSONParser,
			nil,
		},
		{
			DefaultYAML,
			defaultYAMLParser,
			defaultYAMLParser,
			nil,
		},
		{
			// Rgistering defaultJSONParser for already registered format DefaultYAML should return error. However, parser registerd for DefaultYAML still persists from previous registration.
			DefaultYAML,
			defaultJSONParser,
			defaultYAMLParser,
			NewError(InvalidOperation, ERR_ALREADYREGISTERED),
		},
	}

	for _, testCase := range testCases {
		actualError := Register(testCase.RegisterFormat, testCase.RegisterParser)
		parser, _ := parsers.Get(testCase.RegisterFormat.String())
		actualParser := parser.(*Parser)
		if !(actualParser == testCase.ExpectedParser &&
			reflect.DeepEqual(actualError, testCase.ExpectedError)) {
			t.Errorf("Register( %v, %v ) returned Error %v and assigned Parser %#v, while expected Error %v and assigned Parser %#v ", testCase.RegisterFormat, testCase.RegisterParser, actualError, actualParser, testCase.ExpectedError, testCase.ExpectedParser)
		}

	}
}
