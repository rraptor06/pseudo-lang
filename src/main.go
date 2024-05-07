// The main file of the project

package main

import (
	"os"
	"pseudo-lang/plc"
)

func main() {
	os.Exit(plc.Compiler(os.Args))
}
