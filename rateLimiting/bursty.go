package main

import (
	"fmt"
	"time"
)

// rate limit requests with additional 'x' requests per second maintaining
// overall rate limit of 'y' requests per second
//
//	we serve the first 3 immediately because of the burstable rate limiting,
//
// then serve the remaining 2 with ~1sec delays each.
func main() {

	start := time.Now()
	burstyRequests := make(chan time.Time, 5)
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1 * time.Second) {
			burstyLimiter <- t
		}
	}()

	for i := 0; i < 5; i++ {
		burstyRequests <- time.Now()
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request processed at ", req)
	}

	fmt.Println(time.Since(start))

}
