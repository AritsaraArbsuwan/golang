package main

import "fmt"

func main() {
	n, e := fmt.Println("Hollo", "World", 123, 456)
	fmt.Println("number of bytes written :", n)
	fmt.Println("write error encountered :", e)
}
