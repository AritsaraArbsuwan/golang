package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 3000; i++ {
		file, err := os.Create(fmt.Sprintf("%v.txt", i))

		if err != nil {
			return
		}
		defer file.Close()

		for x := 1; x <= 1000; x++ {

			file.WriteString("tanchira \n")
		}
		for x := 1; x <= 1000; x++ {

			file.WriteString("Peawkrathok \n")
		}
		for x := 1; x <= 1000; x++ {

			file.WriteString("15 \n")
		}

	}

	fmt.Println(time.Since(start))
}
