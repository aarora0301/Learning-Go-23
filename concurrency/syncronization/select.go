package main

import (
	"fmt"
	"time"
)

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

	for { // as there are two producers, expect response twice, so select has to run twice
		// if i >2, it will be deadlock
		// use default case to avoid deadlock (i>2)

		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c1:
			fmt.Println("received", msg2)
		default:
			fmt.Println("*")
		}
	}

	fmt.Println("time taken:", time.Since(start))
}
