package main

import "fmt"

func writeLine(a ...interface{}) {
	for _, V := range a {
		fmt.Println(v)
	}

}

func main() {
	writeLine(1, 3.14, "Hello", true)

}
