package main

import "fmt"

func sum(c chan int, numbar ...int) {
	sum := 0
	for _, v := range numbar {
		sum = sum + v
	}
	c <- sum
}

func printer(c chan int) {
	numbar := <-c
	fmt.Println(numbar)
}

func main() {
	c := make(chan int)
}
