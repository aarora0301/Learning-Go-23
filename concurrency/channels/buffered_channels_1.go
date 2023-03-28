package main

import "fmt"

// By default channels are unbuffered, meaning that they will only accept sends (chan <-)
// if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels
// accept a limited number of values without a corresponding receiver for those values.
func main() {

	ch := make(chan string, 2)

	ch <- "ping"
	ch <- "pong"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
