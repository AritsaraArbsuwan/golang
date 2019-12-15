package main

import "fmt"

func suntract(number *int) {
	*number = number
}

func main() {
	x := 10
	suntract(&x)
	fmt.Println(x)
}
