package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

func getnsps(path string, extention string) (files []string) {

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

type NspLinkIndex struct {
	Files       []string
	Directories int
}

func main() {
	files := getnsps(".", ".go")
	fmt.Println(files)

	webpage := "http://192.168.178.7:2480/"

	for i, s := range files {
		files[i] = webpage + url.QueryEscape(s)
		fmt.Println(files)
	}

	indexFile := NspLinkIndex{
		Files:       files,
		Directories: 1,
	}
	var jsonData []byte
	jsonData, err := json.Marshal(indexFile)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}
