package main

import "fmt"
import	m "myMath"
)

func main() {
	avg := m.myMath.Average(15, 16.5, 20.3, 32)
	fmt.Println(avg)
}
