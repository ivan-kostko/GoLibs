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
	"flag"
	"fmt"
	"os"
)

var version = "1.0.1"

var (
    templateFileExtention = "tmpl"
    templatesSubFolder    = "CodeGeneratorTemplates"
)


func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of CodeGenerator:")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "  CodeGenerator -pointer -type=<type> -template=<templateFile>")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "version: ", version)
		flag.PrintDefaults()
	}

	flag.CommandLine.Init("", flag.ExitOnError)
}


func main() {
	typePointer  := flag.Bool("pointer", false, "Determines whether a type is a pointer or not")
	typeName     := flag.String("type", "", "Type fed to CodeGenerator")
	packageName  := flag.String("package", "", "Package name")
    templateName := flag.String("template", "", "Template name. Code Generator looks for templateName.tmpl file at execution folder. IF no file found there then at current working folder.")

	flag.Parse()

	if *typeName == "" || *templateName == "" {
		flag.Usage()
		return
	}

    // Get Template
    tmpl, err := getTemplateFromFile(*templateName)
	if err != nil {
		panic(err)
	}

    *packageName = getPackageName(*packageName)

	outputFullFileName := getOutputFullFileName(*typeName, *templateName)

    writer, err := os.Create(outputFullFileName)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	generator := NewGenerator(tmpl)

	m := getNewMetadata(*typeName, *typePointer, *packageName)
	if err := generator.Generate(writer, m); err != nil {
		panic(err)
	}

	fmt.Printf("Generated file %s by template %s\n", outputFullFileName, *templateName)
}


