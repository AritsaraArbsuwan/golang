package main

func sum(numbres ...int) int {
	total := 0
	for _, n := range numbres {
		total = total + n
	}
	return total
}
