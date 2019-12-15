package main

import "fmt"

func writeLine(a ...interface{}) {
	for _, V := range a {
		fmt.Println(v)
	}

}
