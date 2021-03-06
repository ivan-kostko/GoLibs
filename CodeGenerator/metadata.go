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
    "fmt"
    "unicode"
    "unicode/utf8"
)

type Metadata struct {
	PackageName   string
	TypeName      string
    FullTypeName  string
    PrivTypeName  string
}

// Converts first rune to lower case
func InitToLower(s string) string {
    if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// Metadata private factory
func getNewMetadata(typeName string, pointerType bool, packageName string) (m Metadata) {
	m.TypeName = typeName
	m.PackageName = packageName

	if pointerType {
		m.FullTypeName = fmt.Sprintf("*%s", m.TypeName)
	} else {
		m.FullTypeName = m.TypeName
	}

    m.PrivTypeName = InitToLower(m.TypeName)

	return m
}
