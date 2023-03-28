package main

import (
	"fmt"
	"time"
)

// main is running
// text runs for 1.5 sec
// go routines are spined after text completes so they are not interleaved
// time taken
// Text : 1.5 sec
// main thread sleeps for 1 sec, so goroutines are allowed to run until main thread is sleeping
// Goroutine 1 : 0.5 sec Goroutine 2: 0.5sec concurrently
// total time : 2.5sec (main thread time + sleep time)
func main() {

	start := time.Now()

	text := func(text string) {
		fmt.Println("started :", text)
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 500) // each loop adds a delay of 500ms
			fmt.Println(text, i)
		}
	}

	text("Main")

	go text("Goroutine 1")

	go text("Goroutine 2")

	time.Sleep(time.Second)
	fmt.Println("time taken", time.Since(start))
}
