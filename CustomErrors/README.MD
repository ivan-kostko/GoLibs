# CustomErrors
--
#### Import

    "github.com/ivan-kostko/GoLibs/CustomErrors"

#### Install

    go get github.com/ivan-kostko/GoLibs/CustomErrors

The package implements extended error functionality, allowing slightly better
error handling than golang standart one.

It gives a possibility to define further behaviour based on error type while
message contains better error description.

### Usage

##### type Error

```go
type Error struct {
	Type    ErrorType
	Message string
}
```

Represents custom error as tuple Type + Message.

##### func  NewError

```go
func NewError(typ ErrorType, msg string) *Error
```
Error factory

##### func (Error) Error

```go
func (e Error) Error() string
```
Implementation of standart error interface

##### type ErrorType

```go
type ErrorType int
```

Represents enum of predefined error types

```go
const (
	BasicError ErrorType = iota
	InvalidOperation
	InvalidArgument
	AccessViolation
	Nonsupported
)
```

##### func (ErrorType) String

```go
func (i ErrorType) String() string
```
