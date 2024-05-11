// Package executor: The package containing the executor functions
package executor

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CopyDirectory Copy the content of a directory to another directory
//
/* src: The source directory */
/* dst: The destination directory */
//
// Returns an error if the copy failed
func CopyDirectory(src string, dst string) error {
	filesList, err := ioutil.ReadDir(src)

	if err != nil {
		return err
	}
	os.RemoveAll(dst)
	err = os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}
	for _, file := range filesList {
		srcPath := fmt.Sprintf("%s/%s", src, file.Name())
		dstPath := fmt.Sprintf("%s/%s", dst, file.Name())
		if file.IsDir() {
			err = CopyDirectory(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			content, err := ioutil.ReadFile(srcPath)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(dstPath, content, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
