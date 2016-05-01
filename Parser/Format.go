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

import "fmt"

import . "github.com/ivan-kostko/GoLibs/CustomErrors"

//go:generate stringer -type=Format

const FORMATIOTAOFFSET = 2

// Represents ENUM of supported codecs
type Format int

const (
	DefaultXML Format = iota + FORMATIOTAOFFSET
	DefaultJSON
	DefaultYAML
)

// Represents factory for Format
func GetFormatByString(str string) (Format, *Error) {
	for i := 0; i < len(_Format_index)-1; i++ {
		if str == _Format_name[_Format_index[i]:_Format_index[i+1]] {

			return Format(i + FORMATIOTAOFFSET), nil
		}
	}
	return 0, NewError(Nonsupported, fmt.Sprintf("Parcer: The codec '%s' is not supported", str))
}
