package main

import (
	"fmt"
	"time"
)
 
func printerO(tick, boom <-chan.time.Time) {
    for{
		select{
		case <-tick:
			fmt.Println("tick.")
		}
}