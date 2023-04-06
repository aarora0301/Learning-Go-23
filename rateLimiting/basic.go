package main

import (
	"fmt"
	"time"
)

// basic rate limiter that allows events up to 1 time per second
func main() {

	start := time.Now()
	// load requests into a channel
	requests := make(chan time.Time, 5)
	for i := 1; i <= 5; i++ {
		requests <- time.Now()
	}

	close(requests)

	ticker := time.Tick(time.Second)

	// one request per second
	for req := range requests {
		<-ticker
		fmt.Println("request processed at ", req)
	}

	fmt.Println(time.Since(start))

}
