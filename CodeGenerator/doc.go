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
    go get "github.com/ivan-kostko/GoLibs/CodeGenerator"

    // To define any other location or/and template file extention for templates - use build flags like the following:
    go get "github.com/ivan-kostko/GoLibs/CodeGenerator" -ldflags "-X github.com/ivan-kostko/GoLibs/CodeGenerator/main.templatesMainFolder=YourLocation github.com/ivan-kostko/GoLibs/CodeGenerator/main.templatesSubFolder= "

Description
    The project CodeGenerator represents //go:generate tool to generate _.go files from templates
    It uses standard text/template notation.
    To prevent template differencies it is loking for templates at get project forlder subfolder (default=CodeGeneratorTemplates) files with predefined extention(default=tmpl)
Example
    //go:generate CodeGenerator -pointer -template=container -type=MyType

    It will generate in currently active folder file mytype_container.go with package named as current folder name.
*/
package main
