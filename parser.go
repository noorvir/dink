package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func parseDinkBlocks(r io.Reader) ([]DinkBlock, error) {
	scanner := bufio.NewScanner(r)
	var dinkBlocks []DinkBlock
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "!#dink") {
			dinkBlocks = append(dinkBlocks, DinkBlock{
				LineNumberStart: 0,
				LineNumberEnd:   0,
				Dependencies: []Dependency{
					{
						Type: DependencyTypeFile,
						Link: line,
						Hash: "",
					},
				},
			})
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return dinkBlocks, nil
}

func parseFile(path string) ([]DinkBlock, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	return parseDinkBlocks(file)
}
