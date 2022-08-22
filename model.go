package main

import "fmt"

type DependencyType string

const (
	DependencyTypeFolder    DependencyType = "folder"
	DependencyTypeFile      DependencyType = "file"
	DependencyTypeLine      DependencyType = "line"
	DependencyTypeLineRange DependencyType = "lineRange"
)

type Dependency struct {
	Type DependencyType `json:"type"`
	Link string         `json:"link"`
	Hash string         `json:"hash"`
}

type DinkBlock struct {
	FilePath        string       `json:"filePath"`
	LineNumberStart int          `json:"lineNumberStart"`
	LineNumberEnd   int          `json:"lineNumberEnd"`
	Dependencies    []Dependency `json:"dependencies"`
}

func (d DinkBlock) String() string {
	p := fmt.Sprintf("Location:%s:%d\n", d.FilePath, d.LineNumberStart)
	p += fmt.Sprintf("Dependencies:\n")

	depStr := ""
	for _, dep := range d.Dependencies {
		depStr = fmt.Sprintf("\tType: %s\n", dep.Type)
		depStr += fmt.Sprintf("\tLink: %s\n", dep.Link)
		depStr += fmt.Sprintf("\tHash: %s\n", dep.Hash)
	}

	return p + depStr
}

type SyncedFilesMap map[string][]DinkBlock

func (sfm SyncedFilesMap) String() string {
	str := ""
	for path, blocks := range sfm {
		str += fmt.Sprintf("File: %s:\n", path)
		for _, block := range blocks {
			str += fmt.Sprintf("%s\n", block)
		}
	}
	return str
}
