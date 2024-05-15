// Package convertor: The package containing the convertor functions
package convertor

import (
	"pseudo-lang/parsing"
)

// ConvertCode Convert the code
//
// Returns 0 if the code was converted successfully, 1 otherwise
func ConvertCode() int {
	code := GetCode()

	code.FileContent = parsing.GetFilesContent(parsing.GetSettings().FilesList)
	if code.FileContent == nil {
		return 1
	}
	if GetAllFunctions(code) != 0 || CheckIndentation(code) != 0 || GetAllVariables(code) != 0 {
		return 1
	}
	return analyseCode(code)
}
