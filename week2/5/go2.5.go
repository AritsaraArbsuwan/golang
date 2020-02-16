package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello World", "wold"))
	fmt.Println(strings.Contains("Hello World", "World"))
}
