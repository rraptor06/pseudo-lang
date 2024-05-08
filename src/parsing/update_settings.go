package parsing

import (
	"fmt"
	"os"
)

// UpdateConvertOnly Updates the settings to only convert the program
//
// Returns 0 if everything went well, 1 if an error occured
func UpdateConvertOnly() int {
	settings := GetSettings()

	if settings.Run == true {
		fmt.Fprintf(os.Stderr, "ERROR: You can't only convert the program if you want to run it !\n")
		return 1
	}
	settings.ConvertOnly = true
	return 0
}

// UpdateOutputDir Updates the settings to change the output directory
//
// Returns 0 if everything went well, 1 if an error occured
func UpdateOutputDir(dir string) int {
	_, statError := os.Stat(dir)
	settings := GetSettings()
	var answer string

	if settings.OutputDir != "output" {
		fmt.Printf("The output directory is already set to \"%s\", do you want to overwrite it ? (y/n) ", settings.OutputDir)
		fmt.Scanln(&answer)
		if answer != "y" {
			return 0
		}
	}
	if os.IsNotExist(statError) == false {
		fmt.Printf("The directory \"%s\" already exists, do you want to overwrite it ? (y/n) ", dir)
		fmt.Scanln(&answer)
		if answer != "y" {
			fmt.Printf("The output directory is set to %s.\n", settings.OutputDir)
			return 0
		}
	}
	settings.OutputDir = dir
	return 0
}

// UpdateExecutable Updates the settings to change the executable name
//
// Returns 0 if everything went well, 1 if an error occured
func UpdateExecutable(executable string) int {
	settings := GetSettings()
	var answer string

	if settings.Executable != "main.out" {
		fmt.Printf("The executable name is already set to \"%s\", do you want to overwrite it ? (y/n) ", settings.Executable)
		fmt.Scanln(&answer)
		if answer != "y" {
			return 0
		}
	}
	settings.Executable = executable
	return 0
}

// UpdateRun Updates the settings to run the program
//
// Returns 0 if everything went well, 1 if an error occured
func UpdateRun() int {
	settings := GetSettings()

	if settings.ConvertOnly == true {
		fmt.Fprintf(os.Stderr, "ERROR: You can't run the program if you want to only convert it !\n")
		return 1
	}
	settings.Run = true
	return 0
}
