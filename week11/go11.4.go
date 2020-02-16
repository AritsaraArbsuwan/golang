package main

import (
	"fmt"
	"time"
)

func main(){
	data := 10
	go func(){
		data = 20
	}
}