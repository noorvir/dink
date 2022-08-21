package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

func listMarkdownFiles(dir string) []string {
	var updateList []string

	// Recursively list all the markdown files in the given path
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			updateList = append(updateList, path)
		}
		return nil
	}); err != nil {
		return nil
	}

	return updateList
}

// UpdateCmd updates the snapshot of the mapping from docs and their dependencies.
func UpdateCmd(cmd *cobra.Command, args []string) error {
	// parse all the markdown files in the given path
	// extract the !#dink directives and parse file/folder paths
	// add file/folder paths to updateList
	// throw warning for paths that don't exist
	// follow each path in updateList and hash the contents
	// save a map of the file paths to the contents hash

	fmt.Println(listMarkdownFiles("testdata"))
	return nil
}
