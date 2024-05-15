package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func analyseReturn(function *FunctionStruct, line string, indentationList *[]string) string {
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
	if strings.HasPrefix(line, "retourner:") == false {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid return in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	index := 0
	for index < indentation {
		newLine += "\t"
		index++
	}
	line = strings.ReplaceAll(line, "(adresse)", "&")
	line = strings.ReplaceAll(line, "(valeur)", "*")
	index = 10
	for _, char := range line[index:] {
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
	newLine += "return"
	if line[index:] != "" {
		newLine += " "
	}
	newLine += line[index:] + ";"
	return newLine
}

func analyseFunction(function *FunctionStruct, line string, indentationList *[]string) string {
	var builtinsList []*Instructions = []*Instructions{
		&Instructions{
			Name:     "afficher",
			Function: analyseAfficher,
		},
		&Instructions{
			Name:     "entier",
			Function: analyseEntier,
		},
		&Instructions{
			Name:     "decimal",
			Function: analyseDecimal,
		},
		&Instructions{
			Name:     "texte",
			Function: analyseTexte,
		},
	}
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
	nb := 0
	for nb < indentation {
		newLine += "\t"
		nb++
	}
	line = strings.ReplaceAll(line, "(adresse)", "&")
	line = strings.ReplaceAll(line, "(valeur)", "*")
	for _, instruction := range builtinsList {
		if strings.HasPrefix(line, instruction.Name+"(") {
			newLine += instruction.Function(function, line, indentationList)
			if newLine == "" {
				return ""
			}
			return newLine
		}
	}
	newLine += line + ";"
	return newLine
}
