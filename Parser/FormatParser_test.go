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

	. "github.com/ivan-kostko/GoLibs/CustomErrors"
)

func TestFormatToString(t *testing.T) {

	testCases := []struct {
		Format   FormatParser
		Expected string
	}{
		{
			XMLDefault,
			"XMLDefault",
		},
		{
			JSONDefault,
			"JSONDefault",
		},
		{
			YAMLDefault,
			"YAMLDefault",
		},
	}
	for _, testCase := range testCases {
		codec := testCase.Format
		expected := testCase.Expected

		actual := codec.String()
		if actual != expected {
			t.Errorf("%#v.String() returned %v while expected %v", codec, actual, expected)
		}
	}
}

func TestGetFormatByString(t *testing.T) {

	testCases := []struct {
		FormatName string
		Expected   FormatParser
	}{
		{
			"XMLDefault",
			XMLDefault,
		},
		{
			"JSONDefault",
			JSONDefault,
		},
		{
			"YAMLDefault",
			YAMLDefault,
		},
	}
	for _, testCase := range testCases {
		formatName := testCase.FormatName
		expected := testCase.Expected

		actual, err := GetFormatByString(formatName)
		if actual != expected || err != nil {
			t.Errorf("GetFormatByString(%#v) returned ( %v, %v ) while expected ( %v, %v )", formatName, actual, err, expected, nil)
		}
	}

	// Test non existent codec
	codec := "Nonexistent Format"
	expectedFormat := FormatParser(0)
	expectedErrorType := Nonsupported
	expectedErrorMSg := "Parcer: The codec 'Nonexistent Format' is not supported"

	actual, err := GetFormatByString(codec)
	if actual != expectedFormat ||
		err.Type != expectedErrorType ||
		err.Message != expectedErrorMSg {
		t.Errorf("GetFormatByString(%#v) returned ( %v, %v ) while expected ( %v, %v )", codec, actual, err, expectedFormat, NewError(expectedErrorType, expectedErrorMSg))

	}
}
