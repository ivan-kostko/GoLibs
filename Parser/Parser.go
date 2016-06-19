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
	"io"

	. "github.com/ivan-kostko/GoLibs/CustomErrors"
)

//go:generate CodeGenerator -template=container -type=Parser

// Defines functionality of parser as combination of two functions: Serialize + Deserialize
type Parser interface {
	Deserializer
	Serializer

	// Factory method creating a new Decoder
	NewDecoder(r io.Reader) (Decoder, *Error)

	// Factory method creating a new Encoder
	NewEncoder(w io.Writer) (Encoder, *Error)
}
