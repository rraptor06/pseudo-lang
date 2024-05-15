// Package convertor: The package containing the convertor functions
package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

// getReferenceNumber Get the number of references in a string
//
/* line: The line to analyse */
//
// Returns the number of references
func getReferenceNumber(line string) int {
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

// checkEntier Check if a line is an integer
//
/* function: The function to analyse */
/* line: The line to analyse */
//
// Returns true if the line is an integer, false otherwise
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
	value := getReferenceNumber(variableName)
	variableName = strings.ReplaceAll(variableName, "&", "")
	variableName = strings.ReplaceAll(variableName, "*", "")
	for _, variable := range function.LocalVars {
		if variable.Name == variableName && (variable.VariableType == "int " || value == getReferenceNumber(variable.VariableType)) {
			return true
		}
	}
	return false
}

// checkDecimal Check if a line is a decimal
//
/* function: The function to analyse */
/* line: The line to analyse */
//
// Returns true if the line is a decimal, false otherwise
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
	value := getReferenceNumber(variableName)
	variableName = strings.ReplaceAll(variableName, "&", "")
	variableName = strings.ReplaceAll(variableName, "*", "")
	for _, variable := range function.LocalVars {
		if variable.Name == variableName && (variable.VariableType == "double " || value == getReferenceNumber(variable.VariableType)) {
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

// checkVoid Check if a line is a void
//
/* function: The function to analyse */
/* line: The line to analyse */
//
// Returns true if the line is a void, false otherwise
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

// analyseAfficher Analyse the "afficher" builtin
//
/* function: The function to analyse */
/* line: The line to analyse */
/* indentationList: The list of indentations */
//
// Returns the new line
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

// analyseEntier Analyse the "entier" builtin
//
/* function: The function to analyse */
/* line: The line to analyse */
/* indentationList: The list of indentations */
//
// Returns the new line
func analyseEntier(function *FunctionStruct, line string, indentationList *[]string) string {
	return ""
}

// analyseDecimal Analyse the "decimal" builtin
//
/* function: The function to analyse */
/* line: The line to analyse */
/* indentationList: The list of indentations */
//
// Returns the new line
func analyseDecimal(function *FunctionStruct, line string, indentationList *[]string) string {
	return ""
}

// analyseTexte Analyse the "texte" builtin
//
/* function: The function to analyse */
/* line: The line to analyse */
/* indentationList: The list of indentations */
//
// Returns the new line
func analyseTexte(function *FunctionStruct, line string, indentationList *[]string) string {
	return ""
}
