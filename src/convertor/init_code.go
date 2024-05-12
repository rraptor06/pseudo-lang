// Package convertor: The package containing the convertor functions
package convertor

// VariableStruct The structure containing the properties of a variable
type VariableStruct struct {
	Name         string
	VariableType string
}

// FunctionStruct The structure containing the code and the properties of a function
type FunctionStruct struct {
	Name             string
	Args             []string
	ArgsNames        []string
	LocalVars        []VariableStruct
	Content          []string
	ConvertedContent []string
	Returns          string
}

// CodeStruct The structure containing the code of the program
type CodeStruct struct {
	FileContent   []string
	MainFunction  *FunctionStruct
	FunctionsList []*FunctionStruct
	GlobalVars    []*VariableStruct
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
				LocalVars:        []VariableStruct{},
				Content:          []string{},
				ConvertedContent: []string{},
				Returns:          "int ",
			},
			FunctionsList: []*FunctionStruct{},
			GlobalVars:    []*VariableStruct{},
		}
	}
	return code
}
