package parsing

import (
	"fmt"
	"os"
)

// ParseFlagsAndFiles Parses the flags and the files given to the program
//
// argv: The arguments given to the program
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
		if argv[i][0] == '-' {
			fmt.Fprintf(os.Stderr, "ERROR: Unknown flag %s !\n", argv[i])
			DisplayHelp()
			return 1
		}
		settings.FilesList = append(settings.FilesList, argv[i])
	}
	return 0
}

// ParseArgs Parses the arguments given to the program
//
// argv: The arguments given to the program
//
// Returns 0 if everything went well, 1 if an error occured
func ParseArgs(argv []string) int {
	if len(argv) == 1 {
		fmt.Fprintf(os.Stderr, "ERROR: No argument given !\n")
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
		fmt.Fprintf(os.Stderr, "ERROR: No file given !\n")
		return 1
	}
	return 0
}
