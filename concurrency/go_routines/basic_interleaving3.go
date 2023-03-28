package main

import (
	"fmt"
	"time"
)

// we see the output of the blocking call first, then the output of the two goroutines.
//The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
// there is no guranatee that the goroutines will finish their work before the main thread
func main() {

	//start := time.Now()

	text := func(text string) {
		fmt.Println("started :", text)
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 500) // each loop adds a delay of 500ms
			fmt.Println(text, i)
		}
	}

	go text("Goroutine 1")

	go text("Goroutine 2")

	text("Main")

	time.Sleep(time.Second)

	//fmt.Println("time taken", time.Since(start))
}
