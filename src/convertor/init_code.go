package convertor

// FunctionStruct The structure containing the code and the properties of a function
type FunctionStruct struct {
	name             string
	args             []string
	argsNames        []string
	content          []string
	convertedContent []string
	returns          string
}

// CodeStruct The structure containing the code of the program
type CodeStruct struct {
	fileContent   []string
	mainFunction  FunctionStruct
	functionsList []FunctionStruct
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
