package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func analyseInitVariable(line string, indentationList *[]string) string {
	newLine := "\t"

	indentation := 0
	for IsIndented(line) {
		line = RemoveIndentation(line)
		indentation++
	}
	if indentation > len(*indentationList) {
		fmt.Fprintf(os.Stderr, "%sERROR: Too much indentation in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	} else if indentation < len(*indentationList) {
		*indentationList = (*indentationList)[:len(*indentationList)-1]
		index := 0
		for index < indentation {
			newLine += "\t"
			index++
		}
		newLine += "}\n\t"
	}
	index := strings.Index(line, ":")
	variableType := GetVariablesType(line[:index], true)
	if variableType == "" {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid variable type in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
		return ""
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
				return ""
			}
		}
	}
	value := "0"
	end := strings.Index(line, "<-")
	if end == -1 {
		end = len(line)
	} else {
		value = line[end+2:]
	}
	name := strToArray(line[index+variableLength : end])
	if len(name) != 1 || !IsValidVariableName(name[0]) {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid variable name in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	index = 0
	for _, char := range value {
		if char != ' ' && char != '\t' {
			break
		}
		index++
	}
	value = strings.ReplaceAll(value, "(adresse)", "&")
	value = strings.ReplaceAll(value, "(valeur)", "*")
	nb := 0
	for nb < indentation {
		newLine += "\t"
		nb++
	}
	newLine += variableType + name[0] + " = " + value[index:] + ";"
	return newLine
}

func analyseVariable(function *FunctionStruct, line string, indentationList *[]string) string {
	newLine := "\t"

	indentation := 0
	for IsIndented(line) {
		line = RemoveIndentation(line)
		indentation++
	}
	if indentation > len(*indentationList) {
		fmt.Fprintf(os.Stderr, "%sERROR: Too much indentation in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	} else if indentation < len(*indentationList) {
		*indentationList = (*indentationList)[:len(*indentationList)-1]
		index := 0
		for index < indentation {
			newLine += "\t"
			index++
		}
		newLine += "}\n\t"
	}
	line = strings.ReplaceAll(line, "(adresse)", "&")
	line = strings.ReplaceAll(line, "(valeur)", "*")
	separation := strings.Index(line, "<-")
	if separation == -1 {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid variable assignation \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	name := strToArray(line[:separation])
	variableName := strings.ReplaceAll(name[0], "*", "")
	variableName = strings.ReplaceAll(variableName, "&", "")
	if len(name) != 1 || !IsValidVariableName(variableName) {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid variable name in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	found := false
	for _, variable := range function.LocalVars {
		if variable.Name == variableName && variable.Depth <= indentation {
			found = true
			break
		}
	}
	if !found {
		fmt.Fprintf(os.Stderr, "%sERROR: Variable \"%s\" not found in function \"%s\" !\n%s", constant.ErrorColor, variableName, function.Name, constant.ResetColor)
		return ""
	}
	index := 0
	for _, char := range line[separation+2:] {
		if char != ' ' && char != '\t' {
			break
		}
		index++
	}
	nb := 0
	for nb < indentation {
		newLine += "\t"
		nb++
	}
	newLine += name[0] + " = " + line[separation+2+index:] + ";"
	return newLine
}
