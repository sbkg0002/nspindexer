package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

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
	files := ListAllNsps(path, ".nsp")

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

	fmt.Println("Written index.tfl:\n" + string(jsonData))
}
