package main

import "fmt"

func main() {

	messages := make(chan string) // deadlock

	fmt.Println(<-messages)

	go func() {
		messages <- "ping"
	}()

}
