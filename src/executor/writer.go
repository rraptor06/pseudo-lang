// Package executor: The package containing the executor functions
package executor

import (
	"pseudo-lang/convertor"
)

// WritePrototypes Write the prototypes of the functions in the project.h file
//
/* outputDir: The directory where the project files is */
/* functionsList: The list of functions */
func WritePrototypes(outputDir string, functionsList []*convertor.FunctionStruct) {
	var prototypeList string

	for _, function := range functionsList {
		prototypeList += function.Returns + function.Name + "("
		for index, arg := range function.Args {
			prototypeList += arg + function.ArgsNames[index]
			if index < len(function.Args)-1 {
				prototypeList += ", "
			}
		}
		prototypeList += ");\n"
	}
	ReplaceInFile(GetFilePath(outputDir, "include/project.h"), "// Function prototypes", prototypeList)
}

// GetFunctionText Get the text of a function
//
/* function: The function */
//
// Returns the text of the function
func GetFunctionText(function *convertor.FunctionStruct) string {
	var functionText string

	functionText += function.Returns + function.Name + "("
	for index, arg := range function.Args {
		functionText += arg + function.ArgsNames[index]
		if index < len(function.Args)-1 {
			functionText += ", "
		}
	}
	functionText += ")\n{\n"
	for _, line := range function.ConvertedContent {
		functionText += line + "\n"
	}
	functionText += "}\n"

	return functionText
}

// WriteFunctions Write the functions in the main.c file
//
/* outputDir: The directory where the project files is */
/* functionsList: The list of functions */
/* main: The main function */
func WriteFunctions(outputDir string, functionsList []*convertor.FunctionStruct, main *convertor.FunctionStruct) {
	var functionsText string

	for _, function := range functionsList {
		functionsText += GetFunctionText(function)
		functionsText += "\n"
	}
	functionsText += GetFunctionText(main)
	ReplaceInFile(GetFilePath(outputDir, "src/main.c"), "// Functions", functionsText)
}

// WriteCode Write the code of the program in the output directory
//
/* outputDir: The directory where the project files is */
func WriteCode(outputDir string) {
	code := convertor.GetCode()

	WritePrototypes(outputDir, code.FunctionsList)
	WriteFunctions(outputDir, code.FunctionsList, code.MainFunction)
}
