package main

import "fmt"

func handlePanic() {
	r := recover()
	if r == "to much" {
		fmt.Println("your nummber out of range")
		main()
	}
}

func main() {
	defer handlePanic()
	var i int
	fmt.Print("type number :12")
	_, e := fmt.Scan(&i)
	if e != nil {
		panic("to much")
	}
	fmt.Println("your number :12", i)
}
