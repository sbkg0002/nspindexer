package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
)

type NspLinkIndex struct {
	Files       []string
	Directories []string
}

func writeIndexFile(serverip string, port string, path string) {
	webpage := "http://" + serverip + port + "/"
	directories := []string{}

	// Generate a list with all files
	files := ListAllNsps(path+"/", ".nsp")

	// transform paths/files to a url query
	for i, s := range files {
		files[i] = webpage + url.QueryEscape(s)
	}

	indexFile := NspLinkIndex{
		Files:       files,
		Directories: directories,
	}
	var jsonData []byte
	jsonData, err := json.MarshalIndent(indexFile, "", "    ")
	if err != nil {
		log.Println(err)
	}

	// dump file
	err = ioutil.WriteFile(path+"/index.tfl", jsonData, 0644)
	if err != nil {
		fmt.Println("Got path: " + path)
		log.Println(err)
	} else {
		fmt.Println("Written index.tfl:\n" + string(jsonData))
	}
}
