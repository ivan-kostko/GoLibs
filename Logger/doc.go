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
Description
    Package Logger provides a common interface and wrapper implementation for logging libraries.
    Contains predefined loggers: StderrLogger(prints error log to Std out)
NB:
    For the moment the package is in development(see TODOs).
    Currently, due to possible changes, it is recommended not to use anything from the package,
    except ILogger interface.
Install
    go get "github.com/ivan-kostko/GoLibs/Logger"
Import
    "github.com/ivan-kostko/GoLibs/Logger"
*/
package Logger
