// Package parsing: The package containing the parsing functions
package parsing

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
)

// ParseFlagsAndFiles Parses the flags and the files given to the program
//
/* argv: The arguments given to the program */
//
// Returns 0 if everything went well, 1 if an error occured
func ParseFlagsAndFiles(argv []string) int {
	settings := GetSettings()

	for i := 1; i < len(argv); i++ {
		if argv[i] == "-c" {
			if UpdateConvertOnly() == 1 {
				return 1
			}
			continue
		}
		if argv[i] == "-d" {
			if UpdateOutputDir(argv[i+1]) == 1 {
				return 1
			}
			i++
			continue
		}
		if argv[i] == "-o" {
			if UpdateExecutable(argv[i+1]) == 1 {
				return 1
			}
			i++
			continue
		}
		if argv[i] == "-r" {
			if UpdateRun() == 1 {
				return 1
			}
			continue
		}
		if argv[i] == "-w" {
			if UpdateWarning() == 1 {
				return 1
			}
			continue
		}
		if argv[i][0] == '-' {
			fmt.Fprintf(os.Stderr, "%sERROR: Unknown flag %s !\n%s", constant.ErrorColor, argv[i], constant.ResetColor)
			DisplayHelp()
			return 1
		}
		settings.FilesList = append(settings.FilesList, argv[i])
	}
	return 0
}

// ParseArgs Parses the arguments given to the program
//
/* argv: The arguments given to the program */
//
// Returns 0 if everything went well, 1 if an error occured
func ParseArgs(argv []string) int {
	if len(argv) == 1 {
		fmt.Fprintf(os.Stderr, "%sERROR: No argument given !\n%s", constant.ErrorColor, constant.ResetColor)
		return 1
	}
	if len(argv) == 2 && (argv[1] == "-h" || argv[1] == "--help") {
		DisplayHelp()
		return 1
	}
	if ParseFlagsAndFiles(argv) == 1 {
		return 1
	}
	if len(GetSettings().FilesList) == 0 {
		fmt.Fprintf(os.Stderr, "%sERROR: No file given !\n%s", constant.ErrorColor, constant.ResetColor)
		return 1
	}
	return 0
}
