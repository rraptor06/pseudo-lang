package parsing

// settingsStruct The structure containing the settings of the program
type settingsStruct struct {
	fileList    []string
	convertOnly bool
	outputDir   string
	executable  string
	run         bool
}

// settings The settings of the program
var settings *settingsStruct

// GetSettings Initializes the settings if they are not already and returns them
//
// Returns the settings of the program
func GetSettings() *settingsStruct {
	if settings == nil {
		settings = &settingsStruct{
			convertOnly: false,
			outputDir:   "",
			executable:  "main.out",
			run:         false,
		}
	}
	return settings
}
