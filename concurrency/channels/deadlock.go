package main

import "fmt"

func main() {

	messages := make(chan string) // deadlock

	// receiver is ready but sender is not // trying to receive from an empty channel
	fmt.Println(<-messages)

	go func() {
		messages <- "ping"
	}()

}
