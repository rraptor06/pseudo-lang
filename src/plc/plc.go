package plc

import (
	"fmt"
	"pseudo-lang/parsing"
)

func Compiler(argv []string) int {
	if parsing.ParseArgs(argv) == 1 {
		return 1
	}
	fmt.Print(parsing.GetSettings())
	return 0
}
