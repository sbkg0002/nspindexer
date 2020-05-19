package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"
	"regexp"
	"github.com/radovskyb/watcher"
)

func writeWhileWatchIndexFile(serverip string, port string, basePath string) {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	// Only notify on:
	w.FilterOps(watcher.Rename, watcher.Move, watcher.Write, watcher.Remove)

	// only nsps
	r := regexp.MustCompile(".*")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
				writeIndexFile(serverip, port, basePath)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	fmt.Println("Hier: " + basePath)
	if err := w.Add(basePath); err != nil {
		log.Fatalln(err)
	}

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(basePath); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()

	// Trigger 2 events after watcher started.
	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Create, nil)
		w.TriggerEvent(watcher.Remove, nil)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

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
