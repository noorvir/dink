package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
)

func listMarkdownFiles(dir string) ([]string, error) {
	var updateList []string

	// Recursively list all the markdown files in the given path
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
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
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %v", err)
	}

	return updateList, nil
}

// UpdateCmd updates the snapshot of the mapping from docs and their dependencies.
func UpdateCmd(cmd *cobra.Command, args []string) error {
	// parse all the markdown files in the given path
	// extract the !#dink directives and parse file/folder paths
	// add file/folder paths to updateList
	// throw warning for paths that don't exist
	// follow each path in updateList and hash the contents
	// save a map of the file paths to the contents hash

	paths, err := listMarkdownFiles("testdata")
	if err != nil {
		return fmt.Errorf("failed to list files: %v", err)
	}

	sfm := SyncedFilesMap{}

	for _, path := range paths {
		dinkBlocks, e := parseFile(path)
		if e != nil {
			return fmt.Errorf("failed to parse file: %v", e)
		}
		if len(dinkBlocks) > 0 {
			sfm[path] = dinkBlocks
		}
	}

	fmt.Printf("%s\n", sfm)

	return nil
}
