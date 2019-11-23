package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("Hello Wollo", "o"))
	fmt.Println(strings.LastIndex("Hello world", "O"))
}
