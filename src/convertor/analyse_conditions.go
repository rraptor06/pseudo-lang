package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func analyseElseIf(line string, indentationList *[]string) string {
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
	newLine += "else if (" + line[9:] + ") {"
	*indentationList = append(*indentationList, "if")
	return newLine
}

func analyseIf(line string, indentationList *[]string) string {
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
	newLine += "if (" + line[3:] + ") {"
	*indentationList = append(*indentationList, "if")
	return newLine
}

func analyseElse(line string, indentationList *[]string) string {
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
	newLine += "else (" + line[6:] + ") {"
	*indentationList = append(*indentationList, "else")
	return newLine
}
