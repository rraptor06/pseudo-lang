package convertor

// VariableStruct The structure containing the properties of a variable
type VariableStruct struct {
	name         string
	variableType string
}

// FunctionStruct The structure containing the code and the properties of a function
type FunctionStruct struct {
	name             string
	args             []string
	argsNames        []string
	localVars        []VariableStruct
	content          []string
	convertedContent []string
	returns          string
}

// CodeStruct The structure containing the code of the program
type CodeStruct struct {
	fileContent   []string
	mainFunction  FunctionStruct
	functionsList []FunctionStruct
	globalVars    []VariableStruct
}

// code The decomposed code of the program
var code *CodeStruct

// GetCode Initializes the settings if they are not already and returns them
//
// Returns the settings of the program
func GetCode() *CodeStruct {
	if code == nil {
		code = &CodeStruct{
			fileContent:   []string{},
			mainFunction:  FunctionStruct{},
			functionsList: []FunctionStruct{},
		}
	}
	return code
}
