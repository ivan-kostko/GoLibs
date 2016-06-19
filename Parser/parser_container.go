
//// DO NOT EDIT. It is auto-generated code! ////

package Parser

import (
	. "github.com/ivan-kostko/GoLibs/CustomErrors"
	tsMap "github.com/ivan-kostko/GoLibs/ThreadSafe/Map"
)

// Predefined list of error messages
const (
	ERR_WONTGET     	  = "Parser: Wont get registered instance for given alias"
	ERR_WRONGREGTYPE      = "Parser: The registered instance is not of Parser type"
	ERR_ALREADYREGISTERED = "Parser: There is already registered instance for provided alias. Wont register second time"
)

var initCapacity = 10

// Represents the map of registered string(implementation alias) + Parser implementation
var parsers = tsMap.New(initCapacity)


// Registers implementation by alias
// In case there is already registered Parser instance with same alias it returns ERR_ALREADYREGISTERED
// skipping further registration steps. So, initial registration stays w/o changes.
//
// For.Ex. having registered "SomeImplementation" + SomeParser, registration of "SomeImplementation" + NewParser returns ERR_ALREADYREGISTERED error,
// keeping initial registration (SomeFormat + SomeParser) w/o changes and available for further use.
func Register(implementationAlias string, impl Parser) *Error {
	if _, ok := parsers.Get(implementationAlias); ok {
		return NewError(InvalidOperation, ERR_ALREADYREGISTERED)
	}
	parsers.Set(implementationAlias, impl)
	return nil
}

// Gets implementation by alias
// In case of error returns nil and InvalidOperation error with one of predefined messages:
// ERR_WONTGET
// ERR_WRONGREGTYPE
func Get(implementationAlias string) (impl Parser, err *Error) {
	r, ok := parsers.Get(implementationAlias)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WONTGET)
	}
	impl, ok = r.(Parser)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WRONGREGTYPE)
	}
	return impl, nil
}

