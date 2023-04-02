package main

import (
	"fmt"
	"time"
)

// run ticker for 2 seconds
func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	// go receiver function : ready to receive forever until gets done signal
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	ticker.Stop() // stop ticker after 2 seconds
	done <- true
	fmt.Println("Ticker stopped")
}
