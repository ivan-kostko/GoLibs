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

//// DO NOT EDIT. It is auto-generated code! ////

package Factory

import (
	. "github.com/ivan-kostko/GoLibs/CustomErrors"
	tsMap "github.com/ivan-kostko/GoLibs/ThreadSafe/Map"
)

// Predefined list of error messages
const (
	ERR_WONTGET     	  = "Factory: Wont get registered instance for given alias"
	ERR_WRONGREGTYPE      = "Factory: The registered instance is not of Factory type"
	ERR_ALREADYREGISTERED = "Factory: There is already registered instance for provided alias. Wont register second time"
)

const INIT_CAPACITY = 10

// Represents the map of registered string(implementation alias) + Factory implementation
var factorys = tsMap.New(INIT_CAPACITY)


// Registers implementation by alias
// In case there is already registered Factory instance with same alias it returns ERR_ALREADYREGISTERED
// skipping further registration steps. So, initial registration stays w/o changes.
//
// For.Ex. having registered "SomeImplementation" + SomeFactory, registration of "SomeImplementation" + NewFactory returns ERR_ALREADYREGISTERED error,
// keeping initial registration (SomeFormat + SomeFactory) w/o changes and available for further use.
func Register(implementationAlias string, impl Factory) *Error {
	if _, ok := factorys.Get(implementationAlias); ok {
		return NewError(InvalidOperation, ERR_ALREADYREGISTERED)
	}
	factorys.Set(implementationAlias, impl)
	return nil
}

// Gets parser by implementation alias
// In case of error returns nil and InvalidOperation error with one of predefined messages:
// ERR_WONTGET
// ERR_WRONGREGTYPE
func Get(implementationAlias string) (impl Factory, err *Error) {
	r, ok := factorys.Get(implementationAlias)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WONTGET)
	}
	impl, ok = r.(Factory)
	if !ok {
		return nil, NewError(InvalidOperation, ERR_WRONGREGTYPE)
	}
	return impl, nil
}

