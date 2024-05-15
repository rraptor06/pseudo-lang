package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func analyseWhile(function *FunctionStruct, line string, indentationList *[]string) string {
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
	if strings.HasPrefix(line, "tant que:") == false {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid loop in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	index := 0
	for index < indentation {
		newLine += "\t"
		index++
	}
	line = strings.ReplaceAll(line, "(adresse)", "&")
	line = strings.ReplaceAll(line, "(valeur)", "*")
	index = 9
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
	newLine += "while (" + line[index:] + ") {"
	*indentationList = append(*indentationList, "while")
	return newLine
}
