package main

import (
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for_,fi := range fileInfos {
		filename := fi.Name()
		fmt.Println(filename)
	}
}
