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


package main

import (
	"text/template"
    "io"
)

type Generator struct {
	template *template.Template
}

// Generator factory
func NewGenerator(templateB *template.Template) *Generator {
    return &Generator{template:    template}
}

// Generates output into writer applying metadata to Generator template
func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {

	return g.template.Execute(writer, metadata)
}

func (g *Generator) template() (*template.Template, error) {

	resource, err := ioutil.ReadFile(g.templateFileName)
    if err != nil {
		return nil, err
	}

	tmpl := template.New("template")

    fmt.Println(string(resource))

    return tmpl.Parse(string(resource))
}

