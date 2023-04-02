package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future - tickers are for when you want to do
// something repeatedly at regular intervals.
func main() {

	ticker := time.NewTicker(time.Millisecond * 500)

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}

}
