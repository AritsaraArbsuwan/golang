package main

import "fmt"

func handlePanic(){
	r := recover()
    if r == "to much"{
		fmt.Println("your nummber out of range")
		main()
	}
}

func main (){