// Package executor: The package containing the executor functions
package executor

import (
	"io/ioutil"
	"strings"
)

// ReplaceInFile Replace a word in a file
//
/* filePath: The path of the file to replace the word in */
/* oldWord: The word to replace */
/* newWord: The word to replace with */
//
// Returns an error if there is one
func ReplaceInFile(filePath, oldWord, newWord string) error {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}
	newContent := strings.Replace(string(content), oldWord, newWord, -1)
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetFilePath Get the path of a file
//
/* outputDirectory: The directory where the Makefile is */
//
// Returns the path of the Makefile
func GetFilePath(outputDirectory string, file string) string {
	if outputDirectory[len(outputDirectory)-1] == '/' {
		return outputDirectory + file
	}
	return outputDirectory + "/" + file
}
