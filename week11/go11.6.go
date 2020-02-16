package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(data *int, mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	*data++
	fmt.Println(time.Since(start), "Increment to", *data){
		start := time.Now()
		defer wg.Done()
		defer mutex.Unlock()
		mutex.Lock()
		fmt.Println(time.Since(start), "Data =", *data)
	}
func main(){
	var mutex sync.Mutex
	var wg.sync.WaitGroup
}
}
