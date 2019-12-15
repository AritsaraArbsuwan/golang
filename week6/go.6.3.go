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
	x := sun(1, 3, ,5 ,7 ,9)
	fmt.Println(x)

	y:= sun()
	fmt.Println(y)
	
}
