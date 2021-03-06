# Default
##### Install

    go get "github.com/ivan-kostko/GoLibs/Parser/XML/Default"

##### Import

    "github.com/ivan-kostko/GoLibs/Parser"
    _ "github.com/ivan-kostko/GoLibs/Parser/XML/Default"


##### Description

The Parser/XML/Default package represents wrapper for encoding/xml Go standard package. It adapts standard lib to interfaces defined at github.com/ivan-kostko/GoLibs/Parser and registers underlying parser implementation at github.com/ivan-kostko/GoLibs/Parser.

#### Usage

##### func  Deserialize

```go
func Deserialize(document []byte, dest interface{}) *Error
```

--
##### func  NewDecoder

```go
func NewDecoder(r io.Reader) (parsers.Decoder, *Error)
```

--
##### func  NewEncoder

```go
func NewEncoder(w io.Writer) (parsers.Encoder, *Error)
```

--
##### func  Serialize

```go
func Serialize(in interface{}) ([]byte, *Error)
```

--
##### type Decoder

```go
type Decoder struct {
}
```

Decoder wrapper

--
##### func (*Decoder) Decode

```go
func (this *Decoder) Decode(in interface{}) *Error
```

--
##### type Encoder

```go
type Encoder struct {
}
```

Decoder wrapper

--
##### func (*Encoder) Encode

```go
func (this *Encoder) Encode(in interface{}) *Error
```

--
