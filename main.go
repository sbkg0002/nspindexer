package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
    "path/filepath"
	"io/ioutil"
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
	Directories []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "[(relative)PATH]")
		return
	}
	path := os.Args[1]
	directories := []string{}
    // Generate a list with all files    
    files := getnsps(path, ".nsp")

	webpage := "http://192.168.178.7:2480/"

    // transform paths/files to a url query
	for i, s := range files {
		files[i] = webpage + url.QueryEscape(s)
	}

	indexFile := NspLinkIndex{
		Files:       files,
		Directories: directories,
	}
	var jsonData []byte
	jsonData, err := json.Marshal(indexFile)
	if err != nil {
		log.Println(err)
    }
    // dump file
    _ = ioutil.WriteFile("index.tfl", jsonData, 0644)

	// fmt.Println(string(jsonData))
}
