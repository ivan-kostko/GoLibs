# Implementations
--
#### Install

    go get "github.com/ivan-kostko/GoLibs/Repository/DataSource/Implementations"

#### Import

    "github.com/ivan-kostko/GoLibs/Repository/DataSource/Implementations"

#### Description

    The package represents container for concreet DataSource implementations

### Usage

```go
const (
	ERR_WONTGET           = "Implementations: Wont get registered instance for given alias"
	ERR_WRONGREGTYPE      = "Implementations: The registered instance is not of DataSourceImplementationFactory type"
	ERR_ALREADYREGISTERED = "Implementations: There is already registered instance for provided alias. Wont register second time"
)
```
Predefined list of error messages

```go
const INIT_CAPACITY = 10
```

##### func  Register

```go
func Register(implementationAlias string, impl DataSourceImplementationFactory) *Error
```
Registers implementation by alias In case there is already registered
DataSourceImplementationFactory instance with same alias it returns
ERR_ALREADYREGISTERED skipping further registration steps. So, initial
registration stays w/o changes.

For.Ex. having registered "SomeImplementation" +
SomeDataSourceImplementationFactory, registration of "SomeImplementation" +
NewDataSourceImplementationFactory returns ERR_ALREADYREGISTERED error, keeping
initial registration (SomeFormat + SomeDataSourceImplementationFactory) w/o
changes and available for further use.

##### type DataSourceImplementationFactory

```go
type DataSourceImplementationFactory func(conf interface{}) (*DataSource, *Error)
```

Represents genric DataSource implementation factory function

##### func  Get

```go
func Get(implementationAlias string) (impl DataSourceImplementationFactory, err *Error)
```
Gets parser by implementation alias In case of error returns nil and
InvalidOperation error with one of predefined messages: ERR_WONTGET
ERR_WRONGREGTYPE
