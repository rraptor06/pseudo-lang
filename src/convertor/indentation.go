// Package convertor: The package containing the convertor functions
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

// RemoveIndentation Remove the indentation of a line
//
/* line: The line to remove the indentation */
//
// Returns the line without the indentation
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

// InvalidIndentation Check if the indentation of a line is invalid
//
/* line: The line to check */
//
// Returns true if the indentation is invalid, false otherwise
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

// CheckIndentation Check the indentation of a code
//
/* code: The code to check */
//
// Returns 0 if the indentation is correct, 1 otherwise
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
