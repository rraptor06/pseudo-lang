package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
)

// IsIndented Check if a line is indented
//
/* line: The line to check */
//
// Returns true if the line is indented, false otherwise
func IsIndented(line string) bool {
	space := 0

	for _, char := range line {
		if char != '\t' && char != ' ' {
			return false
		}
		if char == '\t' {
			space += 4
		} else {
			space++
		}
		if space >= 4 {
			break
		}
	}
	return true
}

func RemoveIndentation(line string) string {
	index := 0
	spaces := 0

	for _, char := range line {
		if char != '\t' && char != ' ' {
			break
		}
		if char == '\t' {
			spaces += 4
		} else {
			spaces++
		}
		index++
		if spaces >= 4 {
			break
		}
	}
	return line[index:]
}

func InvalidIndentation(line string) bool {
	space := 0

	for _, char := range line {
		if char != '\t' && char != ' ' {
			break
		}
		if char == '\t' {
			space++
			for space%4 != 0 {
				space++
			}
		} else {
			space++
		}
	}
	if space%4 != 0 {
		return true
	}
	return false
}

func CheckIndentation(code *CodeStruct) int {
	for _, function := range code.FunctionsList {
		for _, line := range function.Content {
			if InvalidIndentation(line) {
				fmt.Fprintf(os.Stderr, "%sERROR: Invalid indentation in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
				return 1
			}
		}
	}
	for _, line := range code.MainFunction.Content {
		if InvalidIndentation(line) {
			fmt.Fprintf(os.Stderr, "%sERROR: Invalid indentation in line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
			return 1
		}
	}
	return 0
}
