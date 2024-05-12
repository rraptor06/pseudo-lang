package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"regexp"
	"strings"
	"unicode"
)

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

func strToArray(str string) []string {
	separator := func(c rune) bool {
		return unicode.IsSpace(c)
	}

	return strings.FieldsFunc(str, separator)
}

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

func GetReturnType(line string) string {
	index := regexp.MustCompile(`->`).FindStringIndex(line)

	if index == nil {
		return "void "
	}
	return GetVariablesType(line[index[1]:], false)
}

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

	NewFunction = &FunctionStruct{
		Name:             functionName,
		Args:             argsList,
		ArgsNames:        argsNamesList,
		LocalVars:        []VariableStruct{},
		Content:          []string{},
		ConvertedContent: []string{},
		Returns:          returnType,
	}
	return NewFunction
}

func GetAllFunctions(code *CodeStruct) int {
	var NewFunction *FunctionStruct

	NewFunction = nil
	for _, line := range code.FileContent {
		if strings.HasPrefix(line, "fonction:") && (line[9] == ' ' || line[9] == '\t') {
			NewFunction = GetNewFunction(line[10:])
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
