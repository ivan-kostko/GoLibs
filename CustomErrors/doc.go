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

/*
Import
    "github.com/ivan-kostko/GoLibs/CustomErrors"

Install
    go get github.com/ivan-kostko/GoLibs/CustomErrors


Description

The package implements extended error functionality, allowing slightly better error handling than golang standart one.

It gives a possibility to define further behaviour based on error type while message contains better error description.

Also, it has an option to report call stack trace ( **recommended only for debugging** ) :

* - being built with `fullStackTrace` tag it will report complete stack trace for NewError/NewErrorF call

*/
package CustomErrors
