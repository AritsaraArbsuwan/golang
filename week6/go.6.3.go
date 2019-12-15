package main

import "fmt"

func sum(numbres ...int) int {
	total := 0
	for _, n := range numbres {
		total = total + n
	}
	return total	
}

func main()  {
	x := sum(1, 3, ,5 ,7 ,9)
	fmt.Println(x)

	y := sum(25 ,0)
	fmt.Println(y)
	
}