# Parser
##### Install

    go get "github.com/ivan-kostko/GoLibs/Parser"

##### Import

    "github.com/ivan-kostko/GoLibs/Parser"

##### Description

    The package provides interface to internal factory of predefined parser wrappers.
    The main purpose of the package is to have standardized parsers for common use.

    TODO(x): Accomplish decoder/encoder functionality

##### Example

    Run with the following parameters:
       go run -ldflags "-X github.com/ivan-kostko/GoLibs/Parser/XML/Default.registerAs=DefaultXML" main.go

    main.go:

       package main

       import (
           "fmt"
           "time"
       	// imports container
       	"github.com/ivan-kostko/GoLibs/Parser"

       	// registers concreet implementation
       	_ "github.com/ivan-kostko/GoLibs/Parser/XML/Default"
       )

       type MyModel struct{
               Name string
               Id    int
               CreatedAt time.Time
           }

       func main() {
       	p, err := Parser.GetParser("DefaultXML")
       	if err != nil {
               fmt.Println(err)
           }
           doc := MyModel{
               "MyName",
               12345,
               time.Now(),
           }

           sdoc, serr := p.Serializer(doc)
           if serr != nil {
               fmt.Println(serr)
           }
           fmt.Println(string(sdoc))

           ddoc := MyModel{}
           derr := p.Deserializer(sdoc, &ddoc)
           if derr != nil {
               fmt.Println(derr)
           }
           fmt.Println(ddoc)

           if doc != ddoc {
               fmt.Println("Serialize-Deserialize returned %v while original model is %v", ddoc, doc)
           }

       }

#### Usage

```go
const (
	ERR_WONTGET           = "Parser: Wont get registered instance for given alias"
	ERR_WRONGREGTYPE      = "Parser: The registered instance is not of Parser type"
	ERR_ALREADYREGISTERED = "Parser: There is already registered instance for provided alias. Wont register second time"
)
```
Predefined list of error messages

--
```go
const (
	ERR_NONSUPPORTED = "Parser : Current parser does not support the method"
)
```
Represents list of standard errors

--
##### func  Register

```go
func Register(implementationAlias string, impl Parser) *Error
```
Registers implementation by alias In case there is already registered Parser instance with same alias it returns ERR_ALREADYREGISTERED skipping further registration steps. So, initial registration stays w/o changes.

For.Ex. having registered "SomeImplementation" + SomeParser, registration of "SomeImplementation" + NewParser returns ERR_ALREADYREGISTERED error, keeping initial registration (SomeFormat + SomeParser) w/o changes and available for further use.

--
##### type Decoder

```go
type Decoder interface {
	// Decode works like [Deserializer](#type-deserializer), except it reads the decoder stream
	Decode(v interface{}) *Error
}
```

A Decoder represents a parser reading a particular input stream

--
##### type Deserializer

```go
type Deserializer interface {
	// Deserializes document into dest according to format
	Deserialize(document []byte, dest interface{}) *Error
}
```

Represents generic de-serialization interface

--
##### type Encoder

```go
type Encoder interface {
	// Encode works like [Serializer](#type-serializer), except it writes into encoder stream
	Encode(v interface{}) *Error
}
```

Encoder represents a parser writing into particular stream

--
##### type Parser

```go
type Parser interface {
	Deserializer
	Serializer

	// Factory method creating a new Decoder
	NewDecoder(r io.Reader) (Decoder, *Error)

	// Factory method creating a new Encoder
	NewEncoder(w io.Writer) (Encoder, *Error)
}
```

Defines functionality of parser as combination of two functions: Serialize + Deserialize

--
##### func  Get

```go
func Get(implementationAlias string) (impl Parser, err *Error)
```
Gets implementation by alias In case of error returns nil and InvalidOperation error with one of predefined messages: ERR_WONTGET ERR_WRONGREGTYPE

--
##### func  NewParser

```go
func NewParser(deserialize func(document []byte, dest interface{}) *Error,
	serialize func(in interface{}) ([]byte, *Error),
	newDecoder func(r io.Reader) (Decoder, *Error),
	newEncoder func(w io.Writer) (Encoder, *Error),
) Parser
```
Parser factory

--
##### type Serializer

```go
type Serializer interface {
	// Serializes in into []byte according to format
	Serialize(in interface{}) ([]byte, *Error)
}
```

Represents generic serialization function.

--
