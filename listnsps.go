package main

import (
	"os"
	"path/filepath"
	"strings"
)

func ListAllNsps(path string, extention string) (files []string) {

	root := path
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != extention {
			return nil
		}
		// Add relative path from root
		files = append(files, strings.Split(path, root)[1])
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
