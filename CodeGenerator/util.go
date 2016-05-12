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

// The file contains helper functions

package main

import (
    "os"
	"path/filepath"
   	"text/template"
    "io/ioutil"
    "fmt"
    "strings"
)

const (
    PNC_WONTGETEXEFOLDER = "generator: Won't get execution folder"
    PNC_WONTGETWRKFOLDER = "generator: Won't get working folder"
    PNC_TMPLFILENOTFOUND = "generator: Template file not found"
)
// The main purpose of the function - to have one standard handler of errors
func onError(err error){
    // it could be replaced with logger or whatever
    fmt.Println(err.Error())
}

// Output accordingly and returns true if error is not nil
func checkOnError(err error) bool {
    isError := (err != nil)
    if isError {
        onError(err)
    }
    return isError
}

//returns binary execution folder
func executionFolder() (string, error) {
    return filepath.Abs(filepath.Dir(os.Args[0]))
}

// returns currently active working folder from which binaries were invoked
func workingFolder() (string, error) {
    return os.Getwd()
}

func getTemplateFromFile(templateName string) (*template.Template, error) {

    templateFileName := templateName+"."+templateFileExtention

    executionFolder, err := executionFolder()
    if checkOnError(err){
        panic (PNC_WONTGETEXEFOLDER)
    }

    fmt.Println(executionFolder)

    templateFullFileName := filepath.Join(executionFolder, templatesSubFolder, templateFileName)

    // Check if template file exists
    if _, err := os.Stat(templateFullFileName); checkOnError(err){
            panic(PNC_TMPLFILENOTFOUND)
        }



    resource, err := ioutil.ReadFile(templateFullFileName)
    if checkOnError(err) {
		panic(err)
	}

	tmpl := template.New(templateName)

    return tmpl.Parse(string(resource))

}

func formatOutputFileName(typeName, templateName string) string {
	return fmt.Sprintf("%s_%s.go", strings.ToLower(typeName), templateName)
}

func getOutputFullFileName(typeName, templateName string) (string) {
    workingFolder, err := workingFolder()
    if checkOnError(err){
        panic(PNC_WONTGETWRKFOLDER)
    }
    return filepath.Join(workingFolder, formatOutputFileName(typeName, templateName))
}

func getPackageName(inputPackageName string) string {
    if inputPackageName == "" {
        workingFolder, err := workingFolder()
        if checkOnError(err) {
    		panic(err)
    	}
        return filepath.Base(workingFolder)
    }
    return inputPackageName
}
