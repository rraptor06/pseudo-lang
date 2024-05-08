package convertor

// functionStruct The structure containing the code and the properties of a function
type functionStruct struct {
	name             string
	args             []string
	content          []string
	convertedContent []string
	returns          string
}

// codeStruct The structure containing the code of the program
type codeStruct struct {
	fileContent   []string
	mainFunction  functionStruct
	functionsList []functionStruct
}

// code The decomposed code of the program
var code *codeStruct

// GetCode Initializes the settings if they are not already and returns them
//
// Returns the settings of the program
func GetCode() *codeStruct {
	if code == nil {
		code = &codeStruct{
			fileContent:   []string{},
			mainFunction:  functionStruct{},
			functionsList: []functionStruct{},
		}
	}
	return code
}
