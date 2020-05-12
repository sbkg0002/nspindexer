package main

// {
//     "files": ["http://192.168.178.7:2480/Super%20Mario%2064%20%28Unofficial%20Native%20Homebrew%20Port%20for%20Nintendo%20Switch%20by%20AtlasNX%29%20%5BNSP%5D/Super%20Mario%2064%20%28Unofficial%20Native%20Homebrew%20Port%20for%20Nintendo%20Switch%20by%20AtlasNX%29%20%5B054507E0B7552000%5D%5Bv0%5D.nsp"],
//     "directories": []
// }

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type nspindex struct {
	Files       []string
	Directories []string
}

func main() {
	webpage := "http://192.168.178.7:2480/"
	nsp := "Super Mario 64 (Unofficial Native Homebrew Port for Nintendo Switch by AtlasNX) [054507E0B7552000][v0].nsp"

	var directories []string
	output := &nspindex{
		Files:       []string{webpage + url.QueryEscape(nsp)},
		Directories: directories,
	}
	outputJ, _ := json.Marshal(output)
	fmt.Println(string(outputJ))

}
