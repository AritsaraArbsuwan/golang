package main

import (
	"sync"
	"time"
)

func increment(data *int, mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
}
