package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func GetVariablesType(line string, inFuction bool) string {
	supplement := " "

	for index := 0; index < len(line); index++ {
		char := line[index]
		if char == ' ' || char == '\t' {
			continue
		}
		if inFuction && strings.HasPrefix(line[index:], "(adresse)") {
			supplement += "&"
			index += 8
			continue
		}
		if (inFuction && strings.HasPrefix(line[index:], "(valeur)")) || (!inFuction && strings.HasPrefix(line[index:], "(adresse)")) {
			supplement += "*"
			index += 7
			if !inFuction {
				index++
			}
			continue
		}
		if strings.TrimRight(line[index:], " \t") == "entier" {
			return "int" + supplement
		}
		if strings.TrimRight(line[index:], " \t") == "decimal" {
			return "double " + supplement
		}
		if strings.TrimRight(line[index:], " \t") == "vide" {
			return "void " + supplement
		}
		break
	}
	return ""
}

func IsValidVariableName(name string) bool {
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') && char != '_' {
			return false
		}
	}
	return true
}

func AddVariable(function *FunctionStruct, line string, lineIndex int) int {
	var variable *VariableStruct

	depth := 0
	for IsIndented(line) {
		line = RemoveIndentation(line)
		depth++
	}
	index := strings.Index(line, ":")
	variableType := GetVariablesType(line[:index], true)
	if variableType == "" {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid variable type in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
		return 1
	}
	index = strings.Index(line, "entier:")
	if index == -1 {
		index = strings.Index(line, "decimal:")
		if index == -1 {
			index = strings.Index(line, "vide:")
			if index == -1 {
				fmt.Fprintf(os.Stderr, "%sERROR: Invalid variable type in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
				return 1
			}
		}
	}
	end := strings.Index(line, "<-")
	if end == -1 {
		end = len(line)
	}
	name := strToArray(line[index+7 : end])
	if len(name) != 1 || !IsValidVariableName(name[0]) {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid variable name in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
		return 1
	}
	variable = &VariableStruct{
		Name:         name[0],
		VariableType: variableType,
		Line:         lineIndex,
		Depth:        depth,
	}
	function.LocalVars = append(function.LocalVars, variable)
	return 0
}

func GetAllVariables(code *CodeStruct) int {
	for _, function := range code.FunctionsList {
		for index, line := range function.Content {
			if strings.Contains(line, ":") &&
				strings.Contains(line, "retourner:") == false && strings.Contains(line, "tant que:") == false &&
				strings.Contains(line, "si:") == false && strings.Contains(line, "sinon:") == false &&
				strings.Contains(line, "sinon si:") == false && AddVariable(code.MainFunction, line, index) == 1 {
				return 1
			}
		}
	}
	for index, line := range code.MainFunction.Content {
		if strings.Contains(line, ":") &&
			strings.Contains(line, "retourner:") == false && strings.Contains(line, "tant que:") == false &&
			strings.Contains(line, "si:") == false && strings.Contains(line, "sinon:") == false &&
			strings.Contains(line, "sinon si:") == false && AddVariable(code.MainFunction, line, index) == 1 {
			return 1
		}
	}
	return 0
}
