package main

import "fmt"

func main() {

	messages := make(chan string)

	messages <- "ping" // deadlock

	fmt.Println(<-messages)
}
