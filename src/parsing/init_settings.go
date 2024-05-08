package parsing

// settingsStruct The structure containing the settings of the program
type settingsStruct struct {
	FilesList   []string
	ConvertOnly bool
	OutputDir   string
	Executable  string
	Run         bool
}

// settings The settings of the program
var settings *settingsStruct

// GetSettings Initializes the settings if they are not already and returns them
//
// Returns the settings of the program
func GetSettings() *settingsStruct {
	if settings == nil {
		settings = &settingsStruct{
			FilesList:   []string{},
			ConvertOnly: false,
			OutputDir:   "",
			Executable:  "main.out",
			Run:         false,
		}
	}
	return settings
}
