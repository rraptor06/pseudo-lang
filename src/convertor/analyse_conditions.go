package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func analyseElseIf(function *FunctionStruct, line string, indentationList *[]string) string {
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
		if len(*indentationList) == 0 || (*indentationList)[len(*indentationList)-1] != "if" {
			fmt.Fprintf(os.Stderr, "%sERROR: Else if without if \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
			return ""
		}
		*indentationList = (*indentationList)[:len(*indentationList)-1]
		index := 0
		for index < indentation {
			newLine += "\t"
			index++
		}
		newLine += "}\n\t"
	} else {
		fmt.Fprintf(os.Stderr, "%sERROR: Else if without if \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	if strings.HasPrefix(line, "sinon si:") == false {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid condition in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
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
	newLine += "else if (" + line[index:] + ") {"
	*indentationList = append(*indentationList, "if")
	return newLine
}

func analyseIf(function *FunctionStruct, line string, indentationList *[]string) string {
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
	if strings.HasPrefix(line, "si:") == false {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid condition in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	index := 0
	for index < indentation {
		newLine += "\t"
		index++
	}
	line = strings.ReplaceAll(line, "(adresse)", "&")
	line = strings.ReplaceAll(line, "(valeur)", "*")
	index = 3
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
	newLine += "if (" + line[index:] + ") {"
	*indentationList = append(*indentationList, "if")
	return newLine
}

func analyseElse(function *FunctionStruct, line string, indentationList *[]string) string {
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
		if len(*indentationList) == 0 || (*indentationList)[len(*indentationList)-1] != "if" {
			fmt.Fprintf(os.Stderr, "%sERROR: Else without if \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
			return ""
		}
		*indentationList = (*indentationList)[:len(*indentationList)-1]
		index := 0
		for index < indentation {
			newLine += "\t"
			index++
		}
		newLine += "}\n\t"
	} else {
		fmt.Fprintf(os.Stderr, "%sERROR: Else without if \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	if strings.HasPrefix(line, "sinon:") == false {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid condition in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return ""
	}
	index := 0
	for index < indentation {
		newLine += "\t"
		index++
	}
	line = strings.ReplaceAll(line, "(adresse)", "&")
	line = strings.ReplaceAll(line, "(valeur)", "*")
	index = 6
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
	newLine += "else (" + line[index:] + ") {"
	*indentationList = append(*indentationList, "else")
	return newLine
}
