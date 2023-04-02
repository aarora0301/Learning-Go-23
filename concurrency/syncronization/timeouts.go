package main

import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	select { // response will be accepted from either of the channels, exits after first response
	case msg := <-c1:
		println(msg)
	case <-time.After(time.Second * 1):
		println("timeout 1") // timeout before go routine is finished
	}

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	select {
	case msg := <-c2: // go routine gets executed
		println(msg)
	case <-time.After(time.Second * 3):
		println("timeout 2")
	}

}
