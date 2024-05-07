// Package parsing : The package containing the parsing functions
package parsing

import "fmt"

func DisplayHelp() {
	fmt.Println("USAGE")
	fmt.Println("\tplc [OPTIONS] [file1.pl file2.pl ...]")
	fmt.Println("DESCRIPTION")
	fmt.Println("\tThe official Pseudo-Lang Compiler")
	fmt.Println("OPTIONS")
	fmt.Println("\t-c\tOnly convert Pseudo-Lang to C, no compilation")
	fmt.Println("\t-d [directory]\tSpecify the code output directory")
	fmt.Println("\t-o [name]\tSpecify the name of the executable (default is main.out)")
	fmt.Println("\t-r\tRun the program after compilation")
}
