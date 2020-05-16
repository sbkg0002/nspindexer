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
		fmt.Println("Usage:", os.Args[0], "IP:PORT")
		fmt.Println("Example:", os.Args[0], "192.168.178.7:2480")
		return
	}
	webpage := "http://" + os.Args[1] +"/"
	directories := []string{}
	// Generate a list with all files
	files := ListAllNsps(".", ".nsp")

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
	err = ioutil.WriteFile("index.tfl", jsonData, 0644)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Written index.tfl:\n" + string(jsonData))
	}
}
