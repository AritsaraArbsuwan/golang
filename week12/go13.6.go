package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFn(path string, info os.FileInfo, err error)
    fmt.Println(path)
