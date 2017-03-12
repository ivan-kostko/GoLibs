# CustomErrors
--
    import "GoLibs/CustomErrors"

### Import

    "github.com/ivan-kostko/GoLibs/CustomErrors"

### Install

    go get github.com/ivan-kostko/GoLibs/CustomErrors


### Description

The package implements extended error functionality, allowing slightly better
error handling than golang standart one.

It gives a possibility to define further behaviour based on error type while
message contains better error description.

Also, it has an option to report call stack trace ( **recommended only for
debugging** ) :

* - being built with `fullStackTrace` tag it will report complete stack trace
for NewError/NewErrorF call

## Usage

```go
const (
	// Represents the limit of considered nested function calls
	MAXNESTEDLEVEL = 128
)
```
Represents implementation limitations constants

#### type Error

```go
type Error struct {
	Type    ErrorType
	Message string
}
```

Represents custom error as tuple Type + Message.

#### func  NewError

```go
func NewError(typ ErrorType, msg string) *Error
```
Error factory

#### func  NewErrorF

```go
func NewErrorF(typ ErrorType, baseMsg string, args ...interface{}) *Error
```
Error factory generating message in fmt.Sprintf manner

#### func (Error) Error

```go
func (e Error) Error() string
```
Implementation of standart error interface

#### type ErrorType

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
List of predefined error types

#### func (ErrorType) String

```go
func (i ErrorType) String() string
```
