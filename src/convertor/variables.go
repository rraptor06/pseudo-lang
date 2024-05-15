// Package convertor: The package containing the convertor functions
package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

// GetVariablesType Get the type of the variables
//
/* line: The line to check */
/* inFunction: True if the line is in a function, false otherwise */
//
// Returns the type of the variables
func GetVariablesType(line string, inFunction bool) string {
	supplement := " "

	for index := 0; index < len(line); index++ {
		char := line[index]
		if char == ' ' || char == '\t' {
			continue
		}
		if inFunction && strings.HasPrefix(line[index:], "(adresse)") {
			supplement += "&"
			index += 8
			continue
		}
		if (inFunction && strings.HasPrefix(line[index:], "(valeur)")) || (!inFunction && strings.HasPrefix(line[index:], "(adresse)")) {
			supplement += "*"
			index += 7
			if !inFunction {
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

// IsValidVariableName Check if a variable name is valid
//
/* name: The name of the variable */
//
// Returns true if the variable name is valid, false otherwise
func IsValidVariableName(name string) bool {
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') && char != '_' {
			return false
		}
	}
	return true
}

// AddVariable Add a variable to the function
//
/* function: The function to add the variable to */
/* line: The line to analyse */
/* lineIndex: The index of the line */
//
// Returns 0 if the variable was added successfully, 1 otherwise
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
	variableLength := 7
	if index == -1 {
		index = strings.Index(line, "decimal:")
		variableLength = 8
		if index == -1 {
			index = strings.Index(line, "vide:")
			variableLength = 5
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
	name := strToArray(line[index+variableLength : end])
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

// GetAllVariables Get all the variables in the code
//
/* code: The code of the program */
//
// Returns 0 if all the variables were added successfully, 1 otherwise
func GetAllVariables(code *CodeStruct) int {
	for _, function := range code.FunctionsList {
		for index, line := range function.Content {
			if strings.Contains(line, ":") &&
				strings.Contains(line, "retourner:") == false && strings.Contains(line, "tant que:") == false &&
				strings.Contains(line, "si:") == false && strings.Contains(line, "sinon:") == false &&
				strings.Contains(line, "sinon si:") == false && AddVariable(function, line, index) == 1 {
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
