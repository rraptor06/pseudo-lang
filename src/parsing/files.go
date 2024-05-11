// Package parsing: The package containing the parsing functions
package parsing

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

// GetFilesContent Check files and get their content
//
/* files: The list of the files to read */
//
// Returns the content of the files given
func GetFilesContent(files []string) []string {
	var filesContent []string

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%sERROR: Could not read the \"%s\" !\n%s", constant.ErrorColor, file, constant.ResetColor)
			return nil
		}
		parts := strings.Split(string(content), "\n")
		for _, part := range parts {
			filesContent = append(filesContent, part)
		}
	}
	return filesContent
}
