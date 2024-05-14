// Package convertor: The package containing the convertor functions
package convertor

import (
	"pseudo-lang/parsing"
)

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
