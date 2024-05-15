package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func getReferencementNumber(line string) int {
	number := 0

	for _, char := range line {
		if char == '&' {
			number--
		}
		if char == '*' {
			number++
		}
	}
	return number
}

func checkEntier(function *FunctionStruct, line string) bool {
	variableName := ""
	end := 0
	number := true

	for _, char := range line {
		if (char < '0' || char > '9') && char != '-' && char != '+' && char != ' ' && char != '\t' {
			number = false
			break
		}
	}
	if number {
		return true
	}
	for _, char := range line {
		if char == ' ' || char == '\t' {
			if end == 1 {
				end = 2
			}
			continue
		}
		if end == 2 {
			return false
		}
		variableName += string(char)
		end = 1
	}
	value := getReferencementNumber(variableName)
	variableName = strings.ReplaceAll(variableName, "&", "")
	variableName = strings.ReplaceAll(variableName, "*", "")
	for _, variable := range function.LocalVars {
		if variable.Name == variableName && (variable.VariableType == "int " || value == getReferencementNumber(variable.VariableType)) {
			return true
		}
	}
	return false
}

func checkDecimal(function *FunctionStruct, line string) bool {
	variableName := ""
	end := 0
	number := true

	for _, char := range line {
		if (char < '0' || char > '9') && char != '-' && char != '+' && char != '.' && char != ' ' && char != '\t' {
			number = false
			break
		}
	}
	if number {
		return true
	}
	for _, char := range line {
		if char == ' ' || char == '\t' {
			if end == 1 {
				end = 2
			}
			continue
		}
		if end == 2 {
			return false
		}
		variableName += string(char)
		end = 1
	}
	value := getReferencementNumber(variableName)
	variableName = strings.ReplaceAll(variableName, "&", "")
	variableName = strings.ReplaceAll(variableName, "*", "")
	for _, variable := range function.LocalVars {
		if variable.Name == variableName && (variable.VariableType == "double " || value == getReferencementNumber(variable.VariableType)) {
			return true
		}
	}
	return false
}

func checkTexte(function *FunctionStruct, line string) bool {
	variableName := ""
	end := 0

	for _, char := range line {
		if char == ' ' || char == '\t' {
			continue
		}
		if end >= 2 {
			return false
		}
		if char == '"' || char == '\'' {
			end++
		}
		if char == '"' || char == '\'' || end == 1 {
			variableName += string(char)
		} else {
			return false
		}
	}
	return true
}

func checkVoid(function *FunctionStruct, line string) bool {
	variableName := ""
	end := 0

	for _, char := range line {
		if char == ' ' || char == '\t' {
			if end == 1 {
				end = 2
			}
			continue
		}
		if end == 2 {
			return false
		}
		variableName += string(char)
		end = 1
	}
	variableName = strings.ReplaceAll(variableName, "&", "")
	variableName = strings.ReplaceAll(variableName, "*", "")
	for _, variable := range function.LocalVars {
		if variable.Name == variableName {
			return true
		}
	}
	return false
}

func analyseAfficher(function *FunctionStruct, line string, indentationList *[]string) string {
	newLine := "my_printf(\""
	var variableList []string

	if line[len(line)-1] != ')' {
		fmt.Fprintf(os.Stderr, "%sERROR: Missing ')' in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	line = line[9 : len(line)-1]
	args := strToArrayInhibitors(line, "+")
	for _, arg := range args {
		if checkEntier(function, arg) {
			newLine += "%d"
			variableList = append(variableList, arg)
			continue
		}
		if checkDecimal(function, arg) {
			newLine += "%f"
			variableList = append(variableList, arg)
			continue
		}
		if checkTexte(function, arg) {
			newLine += "%s"
			variableList = append(variableList, arg)
			continue
		}
		if checkVoid(function, arg) {
			newLine += "%p"
			variableList = append(variableList, arg)
			continue
		}
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid argument for afficher builtin \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	newLine += "\""
	for _, variable := range variableList {
		newLine += ", " + variable
	}
	newLine += ");"
	return newLine
}

func analyseEntier(function *FunctionStruct, line string, indentationList *[]string) string {
	return ""
}

func analyseDecimal(function *FunctionStruct, line string, indentationList *[]string) string {
	return ""
}

func analyseTexte(function *FunctionStruct, line string, indentationList *[]string) string {
	return ""
}
