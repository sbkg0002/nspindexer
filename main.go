package main

import (
	// "log"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func main() {
	files, _ := WalkMatch("/usr/src/", "*.go")
	fmt.Println(files)

	webpage := "http://192.168.178.7:2480/"

	for i, s := range files {
		files[i] = webpage + url.QueryEscape(s)
		fmt.Println(files)

	}
}
