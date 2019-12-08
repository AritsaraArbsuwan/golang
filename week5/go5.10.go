package main

import "fmt"

func main() {
	alphabets := [4]string{"A", "B", "C", "D"}
	fmt.Println(alphabets)

	x := alphabets[0:2]
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := make([]int, 5, 10)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

}
