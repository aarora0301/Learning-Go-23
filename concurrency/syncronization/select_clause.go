package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations.
// total time is 1 sec and not 2sec, bcz go routines execute concurrently
func main() {

	start := time.Now()
	c1 := make(chan string)
	//c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ { // as there are two producers, expect response twice, so select has to run twice
		// if i>2, it will be deadlock
		// if i<1 , only one response will be received
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c1:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("time taken:", time.Since(start))

}
