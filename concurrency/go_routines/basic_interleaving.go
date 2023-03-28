package main

import (
	"fmt"
	"time"
)

// only main routine is running even if there is a sleep of 500 ms in each loop still it is not suspended
func main() {

	start := time.Now()

	text := func(text string) {
		fmt.Println("started :", text)
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 500) //
			fmt.Println(text, i)
		}
	}

	text("Main")

	go text("Goroutine 1")

	go text("Goroutine 2")

	fmt.Println("time taken", time.Since(start))
}
