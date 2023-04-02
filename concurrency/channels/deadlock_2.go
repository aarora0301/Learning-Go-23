package main

import "fmt"

func main() {

	messages := make(chan string)

	messages <- "ping" // deadlock

	// sender is ready but receiver is not // trying to send to a full channel
	fmt.Println(<-messages)
}
