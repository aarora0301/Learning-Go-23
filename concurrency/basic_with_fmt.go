package main

import (
	"fmt"
)

func main() {

	text := func(text string) {
		fmt.Println("started :", text)
		for i := 0; i < 3; i++ {
			fmt.Println(text, i)
		}
	}

	text("Main")

	go text("Goroutine 1")

	go text("Goroutine 2")

	fmt.Println("Blocking call, still main thread is running")

}
