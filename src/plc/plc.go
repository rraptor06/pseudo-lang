package plc

import (
	"pseudo-lang/executor"
	"pseudo-lang/parsing"
)

func Compiler(argv []string) int {
	if parsing.ParseArgs(argv) == 1 {
		return 1
	}
	parsing.GetFilesContent(parsing.GetSettings().FilesList)
	return executor.LaunchProgram()
}
