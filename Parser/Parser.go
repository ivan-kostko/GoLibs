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

// Represents list of standard errors
const (
	ERR_NONSUPPORTED = "Parser : Current parser does not support the method"
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

// Represents parser implementation
type parser struct {
	deserialize func(document []byte, dest interface{}) *Error
	serialize   func(in interface{}) ([]byte, *Error)
	newDecoder  func(r io.Reader) (Decoder, *Error)
	newEncoder  func(w io.Writer) (Encoder, *Error)
}

func (this *parser) Deserialize(document []byte, dest interface{}) *Error {
	if this.deserialize == nil {
		return NewError(Nonsupported, ERR_NONSUPPORTED)
	}
	return this.deserialize(document, dest)
}

func (this *parser) Serialize(in interface{}) ([]byte, *Error) {
	if this.serialize == nil {
		return nil, NewError(Nonsupported, ERR_NONSUPPORTED)
	}
	return this.serialize(in)
}

func (this *parser) NewDecoder(r io.Reader) (Decoder, *Error) {
	if this.newDecoder == nil {
		return nil, NewError(Nonsupported, ERR_NONSUPPORTED)
	}
	return this.newDecoder(r)
}
func (this *parser) NewEncoder(w io.Writer) (Encoder, *Error) {
	if this.newEncoder == nil {
		return nil, NewError(Nonsupported, ERR_NONSUPPORTED)
	}
	return this.newEncoder(w)
}

// Parser factory
func NewParser(deserialize func(document []byte, dest interface{}) *Error,
	serialize func(in interface{}) ([]byte, *Error),
	newDecoder func(r io.Reader) (Decoder, *Error),
	newEncoder func(w io.Writer) (Encoder, *Error),
) Parser {
	return &parser{
		deserialize: deserialize,
		serialize:   serialize,
		newDecoder:  newDecoder,
		newEncoder:  newEncoder,
	}
}
