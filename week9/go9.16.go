package main

import "fmt"

type I interface{}

func main() {
	var i Ii = "Hello"
	i = "Hello"
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok:= i.(float64)
	fmt.Println(f, ok)

	b :=;.(bool)
	fmt.Println(b)
}
