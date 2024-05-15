// Package convertor: The package containing the convertor functions
package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"regexp"
	"strings"
	"unicode"
)

// GetFunctionName Get the name of a function
//
/* line: The line to check */
//
// Returns the name of the function
func GetFunctionName(line string) string {
	var name string
	state := 0

	name = ""
	for _, char := range line {
		if state == 0 && (char == ' ' || char == '\t') {
			continue
		}
		state = 1
		if state == 1 && char == '(' {
			break
		}
		if state == 1 && ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '_') {
			name += string(char)
			continue
		}
		if char == ' ' || char == '\t' {
			state = 2
			continue
		}
		return ""
	}
	return name
}

// strToArray Convert a string to an array
//
/* str: The string to convert */
//
// Returns the array
func strToArray(str string) []string {
	separator := func(c rune) bool {
		return unicode.IsSpace(c)
	}

	return strings.FieldsFunc(str, separator)
}

// strToArrayInhibitors Convert a string to an array with inhibitors
//
/* str: The string to convert */
/* separators: The separators */
//
// Returns the array
func strToArrayInhibitors(str string, separators string) []string {
	var array []string
	newStr := ""
	find := false
	inhibitor := 0

	for _, char := range str {
		find = false
		if inhibitor == 0 {
			for _, separator := range separators {
				if char == separator {
					if newStr != "" {
						array = append(array, newStr)
						newStr = ""
					}
					find = true
					break
				}
			}
		}
		if !find {
			if inhibitor != 0 && char == int32(inhibitor) {
				inhibitor = 0
			} else if inhibitor == 0 && (char == '"' || char == '\'') {
				inhibitor = int(char)
			}
			newStr += string(char)
		}
	}
	if newStr != "" {
		array = append(array, newStr)
	}
	return array
}

// GetFunctionArgs Get the arguments of a function
//
/* line: The line to check */
//
// Returns the list of arguments,
func GetFunctionArgs(line string) ([]string, []string, int) {
	var argsList []string
	var argsNamesList []string

	for index := len(line) - 1; index >= 0; index-- {
		if line[index] == ' ' || line[index] == '\t' {
			line = line[:index]
		} else if line[index] == ')' {
			line = line[:index]
			break
		} else {
			return nil, nil, 1
		}
	}
	for _, arg := range strings.Split(line, ",") {
		decomposedArg := strToArray(arg)
		if len(decomposedArg) != 2 {
			return nil, nil, 1
		}
		if GetVariablesType(decomposedArg[0], false) == "" || !IsValidVariableName(decomposedArg[1]) {
			return nil, nil, 1
		}
		argsList = append(argsList, GetVariablesType(decomposedArg[0], false))
		argsNamesList = append(argsNamesList, decomposedArg[1])
	}
	return argsList, argsNamesList, 0
}

// GetReturnType Get the return type of a function
//
/* line: The line to check */
//
// Returns the return type
func GetReturnType(line string) string {
	index := regexp.MustCompile(`->`).FindStringIndex(line)

	if index == nil {
		return "void "
	}
	return GetVariablesType(line[index[1]:], false)
}

// GetNewFunction Get and create a new function
//
/* line: The line to check */
//
// Returns the new function
func GetNewFunction(line string) *FunctionStruct {
	var NewFunction *FunctionStruct
	functionName := GetFunctionName(line)
	if functionName == "" {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid function name !\n%s", constant.ErrorColor, constant.ResetColor)
		return nil
	}
	indexStart := regexp.MustCompile(`\(`).FindStringIndex(line)
	if indexStart == nil {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid function declaration !\n%s", constant.ErrorColor, constant.ResetColor)
		return nil
	}
	indexEnd := regexp.MustCompile(`->`).FindStringIndex(line)
	if indexEnd == nil {
		indexEnd = []int{len(line)}
	}
	argsList, argsNamesList, error := GetFunctionArgs(line[indexStart[1]:indexEnd[0]])
	if error != 0 {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid function arguments !\n%s", constant.ErrorColor, constant.ResetColor)
		return nil
	}
	returnType := GetReturnType(line)
	if returnType == "" {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid return type !\n%s", constant.ErrorColor, constant.ResetColor)
		return nil
	}
	var variableStructList []*VariableStruct
	for index, arg := range argsList {
		variableStructList = append(variableStructList, &VariableStruct{
			Name:         argsNamesList[index],
			VariableType: arg,
			Line:         0,
			Depth:        0,
		})
	}

	NewFunction = &FunctionStruct{
		Name:             functionName,
		Args:             argsList,
		ArgsNames:        argsNamesList,
		LocalVars:        variableStructList,
		Content:          []string{},
		ConvertedContent: []string{},
		Returns:          returnType,
	}
	return NewFunction
}

// GetAllFunctions Get all the functions in the code
//
/* code: The code of the program */
//
// Returns 0 if all the functions were added successfully, 1 otherwise
func GetAllFunctions(code *CodeStruct) int {
	var NewFunction *FunctionStruct

	NewFunction = nil
	for _, line := range code.FileContent {
		if strings.HasPrefix(line, "fonction:") {
			NewFunction = GetNewFunction(line[9:])
			if NewFunction == nil {
				return 1
			}
			code.FunctionsList = append(code.FunctionsList, NewFunction)
			continue
		}
		if NewFunction != nil && IsIndented(line) {
			NewFunction.Content = append(NewFunction.Content, RemoveIndentation(line))
			continue
		} else {
			NewFunction = nil
		}
		code.MainFunction.Content = append(code.MainFunction.Content, line)
	}
	return 0
}
