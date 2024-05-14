// Package parsing: The package containing the parsing functions
package parsing

import "fmt"

// DisplayHelp Displays the help message
func DisplayHelp() {
	fmt.Println("USAGE")
	fmt.Println("\tplc [OPTIONS] [file1.psl file2.psl ...]")
	fmt.Println("DESCRIPTION")
	fmt.Println("\tThe official Pseudo-Lang Compiler")
	fmt.Println("OPTIONS")
	fmt.Println("\t-c\t\tOnly convert Pseudo-Lang to C, no compilation")
	fmt.Println("\t-d [directory]\tSpecify the code output directory (default is output)")
	fmt.Println("\t-o [name]\tSpecify the name of the executable (default is main.out)")
	fmt.Println("\t-r\t\tRun the program after compilation")
	fmt.Println("\t-w\t\tDisable warning display")
}
