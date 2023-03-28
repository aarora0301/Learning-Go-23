package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	text := func(text string) {
		fmt.Println("started :", text)
		for i := 0; i < 3; i++ {
			fmt.Println(text, i)
		}

	}

	text("Main")

	go text("Goroutine 1")

	go text("Goroutine 2")

	time.Sleep(time.Second) // main thread sleeps for one second to allow other go routines to finish
	fmt.Println("time taken", time.Since(start))
}
