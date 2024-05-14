package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

func analyseFunction(line string, indentation *int) string {
	return ""
}

func analyseReturn(line string, indentation *int) string {
	new_line := ""

	currentIndentation := 0
	for IsIndented(line) {
		line = RemoveIndentation(line)
		currentIndentation++
	}
	if currentIndentation > *indentation {
		fmt.Fprintf(os.Stderr, "%sERROR: Too much indentation in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return new_line
	}
	*indentation = currentIndentation
	if strings.HasPrefix(line, "retourner:") == false {
		fmt.Fprintf(os.Stderr, "%sERROR: Invalid return in line \"%s\" !%s\n", constant.ErrorColor, line, constant.ResetColor)
		return new_line
	}
	new_line = "return" + line[10:] + ";"
	return new_line
}
