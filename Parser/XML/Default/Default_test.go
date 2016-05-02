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
	"testing"

	parsers "github.com/ivan-kostko/GoLibs/Parser"
)

func TestInitRegistration(t *testing.T) {

	p, err := parsers.GetParserByFormat(parsers.DefaultXML)
	if err != nil {
		t.Errorf("parser.GetParserByFormat(parser.DefaultXML) returned error %v while no error expected", err)
	}
	expectedParser := &parser
	if p != expectedParser {
		t.Errorf("parser.GetParserByFormat(parser.DefaultXML) returned parser %v while expected %v", *p, expectedParser)
	}
}
