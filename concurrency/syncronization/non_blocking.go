package main

import "fmt"

// Basic sends and receives on channels are blocking. However, we can use select with a
// default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	// non-blocking receive, when nothing is sent to channel
	select {
	case msg1 := <-chan1:
		fmt.Println("received", msg1)
	default:
		fmt.Println("no message received")
	}

	// non-blocking send when receiver is not ready
	select {
	case chan2 <- "ping":
		fmt.Println("message sent")
	default:
		fmt.Println("no message sent")
	}

	// non-blocking forever when receiver is not ready
OUTER:
	for {
		select {
		case chan2 <- "ping":
			fmt.Println("message sent")
		default:
			fmt.Println("no message sent")
			break OUTER
		}

	}

	// non-blocking forever when receiver is ready
	// i. sender sends continuously
	// ii. receiver receives only once

	go func() {
		<-chan2 // receiver is ready // message is received only once
	}()

OUTER1:
	for {
		select {
		case chan2 <- "ping": // sender sends multiple times
			fmt.Println("message sent")
		default:
			fmt.Println("no message sent")
			break OUTER1
		}

	}

	// non-blocking forever when receiver is ready
	// i. sender sends continuously
	// ii. receiver receives continuously

	go func() {

		for {
			<-chan2 // receiver is continuously receiving
		}

	}()

OUTER2:
	for {
		select {
		case chan2 <- "ping": // sender sends multiple times : infinite for loop
			fmt.Println("message sent")
		default:
			fmt.Println("no message sent")
			break OUTER2
		}

	}

}
