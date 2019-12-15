package main

import "fmt"

func say() {
	fmt.Scanln("Hello")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	say()
	greet("Doremon")
}
