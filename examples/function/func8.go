// Go in action
// @jeffotoni
// 2019-01-24

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Listing our example directory
// Listing our example directory recursively
func main() {

	// Capturing our path that is in the environment
	gopath := os.Getenv("PWD")

	// directory we want to list
	gopath += "/examples"

	// Making call in function
	list := ListDir(gopath)

	// listing the function return
	for i, p := range list {
		fmt.Printf("[%d:%s===%s]\n", i, path.Dir(p), path.Base(p))
	}
}

// This function uses pkg filepath.Walk, it is
// a recursive function, where it will go through
// our directory and its subfolders.
func ListDir(rootpath string) []string {

	list := make([]string, 0)

	// recursive call
	// This function receives a function as parameter and after going through all levels it ends.
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".git" && filepath.Ext(path) != ".svn" {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}
