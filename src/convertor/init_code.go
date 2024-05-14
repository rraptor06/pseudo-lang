// Package convertor: The package containing the convertor functions
package convertor

// VariableStruct The structure containing the properties of a variable
type VariableStruct struct {
	Name         string
	VariableType string
	Line         int
	Depth        int
}

// FunctionStruct The structure containing the code and the properties of a function
type FunctionStruct struct {
	Name             string
	Args             []string
	ArgsNames        []string
	LocalVars        []*VariableStruct
	Content          []string
	ConvertedContent []string
	Returns          string
}

// CodeStruct The structure containing the code of the program
type CodeStruct struct {
	FileContent   []string
	MainFunction  *FunctionStruct
	FunctionsList []*FunctionStruct
}

// code The decomposed code of the program
var code *CodeStruct

// GetCode Initializes the settings if they are not already and returns them
//
// Returns the settings of the program
func GetCode() *CodeStruct {
	if code == nil {
		code = &CodeStruct{
			FileContent: []string{},
			MainFunction: &FunctionStruct{
				Name: "main",
				Args: []string{
					"int ",
					"char **",
					"char **",
				},
				ArgsNames: []string{
					"argc",
					"argv",
					"envp",
				},
				LocalVars: []*VariableStruct{
					&VariableStruct{
						Name:         "argc",
						VariableType: "int ",
						Line:         0,
						Depth:        0,
					},
					&VariableStruct{
						Name:         "argv",
						VariableType: "char **",
						Line:         0,
						Depth:        0,
					},
					&VariableStruct{
						Name:         "envp",
						VariableType: "char **",
						Line:         0,
						Depth:        0,
					},
				},
				Content:          []string{},
				ConvertedContent: []string{},
				Returns:          "int ",
			},
			FunctionsList: []*FunctionStruct{},
		}
	}
	return code
}
