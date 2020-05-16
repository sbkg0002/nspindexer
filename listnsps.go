package main

import (
	"os"
	"path/filepath"
)

func ListAllNsps(path string, extention string) (files []string) {

	root := path
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != extention {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
