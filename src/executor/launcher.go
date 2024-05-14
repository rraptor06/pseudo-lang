// Package executor: The package containing the executor functions
package executor

import (
	"fmt"
	"os"
	"os/exec"
	"pseudo-lang/constant"
	"pseudo-lang/parsing"
)

// LaunchProgram Create the project and run it
//
// Returns 0 if the program ran successfully, 1 otherwise
func LaunchProgram() int {
	settings := parsing.GetSettings()

	if CopyDirectory("template", settings.OutputDir) != nil {
		fmt.Fprintf(os.Stderr, "%sERROR: Can't create the output directory !\n%s", constant.ErrorColor, constant.ResetColor)
		return 1
	}
	ReplaceInFile(GetFilePath(settings.OutputDir, "Makefile"), "plc-project", settings.Executable)
	WriteCode(settings.OutputDir)
	if settings.ConvertOnly {
		fmt.Printf("%sThe program has been converted in the \"%s\" directory.\n%s", constant.SuccessColor, settings.OutputDir, constant.ResetColor)
		return 0
	}
	cmd := exec.Command("make")
	cmd.Dir = settings.OutputDir
	if settings.Warning == true {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sERROR: Couldn't compile the program !\n%s", constant.ErrorColor, constant.ResetColor)
		return 1
	} else if settings.Warning == false {
		fmt.Printf("%s%s compiled.\n%s", constant.CompilationColor, settings.Executable, constant.ResetColor)
	}
	if settings.Run {
		fmt.Printf("%sRunning the program...\n%s", constant.SuccessColor, constant.ResetColor)
		cmd := exec.Command("./" + settings.Executable)
		cmd.Dir = settings.OutputDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	return 0
}
