package main

func Average(numbers...float64) float64 {
	var total float64
	for_, v := range numbers {
		total = total + v
	}
	return total / float64(len(numbers))
