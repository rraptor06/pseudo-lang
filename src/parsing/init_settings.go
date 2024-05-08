package parsing

// SettingsStruct The structure containing the settings of the program
type SettingsStruct struct {
	FilesList   []string
	ConvertOnly bool
	OutputDir   string
	Executable  string
	Run         bool
}

// settings The settings of the program
var settings *SettingsStruct

// GetSettings Initializes the settings if they are not already and returns them
//
// Returns the settings of the program
func GetSettings() *SettingsStruct {
	if settings == nil {
		settings = &SettingsStruct{
			FilesList:   []string{},
			ConvertOnly: false,
			OutputDir:   "output",
			Executable:  "main.out",
			Run:         false,
		}
	}
	return settings
}
