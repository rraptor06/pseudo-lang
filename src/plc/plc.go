// Package plc: The package containing the compiler functions
package plc

import (
	"pseudo-lang/convertor"
	"pseudo-lang/executor"
	"pseudo-lang/parsing"
)

// Compiler Launch the compiler
//
/* argv: The arguments of the program */
//
// Returns 0 if the program ran successfully, 1 otherwise
func Compiler(argv []string) int {
	if parsing.ParseArgs(argv) == 1 {
		return 1
	}
	if convertor.ConvertCode() == 1 {
		return 1
	}
	return executor.LaunchProgram()
}
