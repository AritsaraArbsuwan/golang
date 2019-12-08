package main

import "fmt"

func main() {
	humans := []string{"Arit", "Boat", "Doremon"}
	names := []string{"Doreme", "Dum"}
	names = append(names, humans...)
	fmt.Println(names)

}
