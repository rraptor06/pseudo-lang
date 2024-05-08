package executor

import (
	"fmt"
	"os"
	"os/exec"
	"pseudo-lang/parsing"
)

func LaunchProgram() int {
	settings := parsing.GetSettings()

	//Update the Makefile

	WriteCode(settings.OutputDir)
	if settings.ConvertOnly {
		return 0
	}
	cmd := exec.Command("make")
	cmd.Dir = settings.OutputDir
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Couldn't compile the program !\n")
		return 1
	}
	if settings.Run {
		exec.Command("./" + settings.Executable)
	}
	return 0
}
