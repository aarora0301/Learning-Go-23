package main

import (
	"fmt"
)

// it’s possible to read from a closed channel
func main() {

	// reading from a buffered channel
	buffer := make(chan string, 2)
	buffer <- "ping"
	buffer <- "pong"

	close(buffer)

	for elem := range buffer {
		fmt.Println(elem)

	}

	// reading from a non buffered channel
	buffer1 := make(chan string)

	go func() {
		for elem := range buffer1 {
			fmt.Println(elem)
		}
	}()

	buffer1 <- "ping"
	close(buffer1)

}
