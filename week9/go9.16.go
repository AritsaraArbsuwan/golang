package main

import "fmt"

type I interface{}

func main() {
	var i Ii = "Hello"
	i = "Hello"
	s, ok := i.(string)
	fmt.Println(s, ok)
}
