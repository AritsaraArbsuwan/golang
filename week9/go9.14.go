package main

import "fmt"

func hello(t interface{}) {
	fmt.Printf("Hello %T \n", t)
}
func hi(t ...interface{}) {
	fmt.Printf("Hi ")
	for _, v := range t {
		fmt.Println(" %T,", v)
	}
	fmt.Println()
}

func main() {
	hello(3.14159265359)
	hello(true)
	hello("Goku")
	hi("Gohan", false, 10, 10e15)

}
