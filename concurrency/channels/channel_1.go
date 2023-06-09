package main

import "fmt"

// channel receive and send are blocking calls
// pipes that connect concurrent goroutines
func main() {

	messages := make(chan string)

	go func() {
		for {
			messages <- "ping"

		}
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
