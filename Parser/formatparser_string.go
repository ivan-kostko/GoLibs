// Code generated by "stringer -type=FormatParser"; DO NOT EDIT

package Parser

import "fmt"

const _FormatParser_name = "XMLDefaultJSONDefaultYAMLDefault"

var _FormatParser_index = [...]uint8{0, 10, 21, 32}

func (i FormatParser) String() string {
	i -= 2
	if i < 0 || i >= FormatParser(len(_FormatParser_index)-1) {
		return fmt.Sprintf("FormatParser(%d)", i+2)
	}
	return _FormatParser_name[_FormatParser_index[i]:_FormatParser_index[i+1]]
}
