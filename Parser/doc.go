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
Install
    go get "github.com/ivan-kostko/GoLibs/Parser"
Import
    "github.com/ivan-kostko/GoLibs/Parser"
Description
    The package provides interface to internal factory of predefined parser wrappers.
    The main purpose of the package is to have standardized parsers for common use.

    TODO(x): Accomplish decoder/encoder functionality
Example
 Run with the following parameters:
    go run -ldflags "-X github.com/ivan-kostko/GoLibs/Parser/XML/Default.RegisterAs=DefaultXML" main.go

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


*/
package Parser
