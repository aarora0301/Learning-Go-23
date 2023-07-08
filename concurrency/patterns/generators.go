package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boringG(msg string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
func main() {

	// Generator Pattern : channel as first class citizen (returning channels)
	c := boringG("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")

	// Channels as a handle on a service
	joe := boringG("Joe")
	anne := boringG("Anne")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-anne)
	}
	fmt.Println("You're both boring; I'm leaving.")

	// Multiplexing : Fan In
	t := fanIN(boringG("Joe"), boringG("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-t)
	}
	fmt.Println("You're both still boring; I'm leaving.")

	// Fan In using select
	q := fanINWithSelect(boringG("Joe"), boringG("Anne"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-q)
	}
	fmt.Println("You're both still boring; I'm leaving!!!!.")

	// Stop after some time is elapsed
	// Add timeout per message and for the entire conversation :
	z := boringG("Anne")
	timeout := time.After(5 * time.Second) // //Timeout for whole conversation using select
	stopAfterTimeElapsed(z, timeout)
	fmt.Println("Timed Out")

	// Stop after a signal is sent
	n := make(chan int)
	stop1 := make(chan bool)
	stopWhenSignalReceived(n, stop1)
	for i := 0; i < 5; i++ {
		n <- i
	}
	stop1 <- true
	fmt.Println("Received Stop signal")

	// Do something when time is elapsed
	m := make(chan int)
	sstop := make(chan bool)
	DoSomethingWhenStopSignalReceived(m, sstop)
	for i := 0; i < 5; i++ {
		m <- i
	}
	sstop <- true
	time.Sleep(1 * time.Second)
	fmt.Println("Did something")
}

func fanIN(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-ch1
		}
	}()
	go func() {
		for {
			c <- <-ch2
		}
	}()
	return c
}

func fanINWithSelect(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				c <- s
			case s := <-ch2:
				c <- s
			}

		}
	}()
	return c
}

func stopAfterTimeElapsed(z chan string, timeout <-chan time.Time) {
	for {
		select {
		case s := <-z:
			fmt.Println(s)
		case <-time.After(100 * time.Millisecond): // act as a timeout for each message
			fmt.Println("You are slow")
			return
		case <-timeout: // act as a timeout for entire conversation
			fmt.Println("You are still slow")
			return
		}
	}
}

func stopWhenSignalReceived(n chan int, stop1 chan bool) {
	go func() {
		for {
			select {
			case temp := <-n:
				fmt.Println("receive from channel", temp)
			case <-stop1:
				fmt.Println("gracefully shutdown")
				return
			}
		}
	}()
}

func DoSomethingWhenStopSignalReceived(n chan int, stop1 chan bool) {
	go func() {
		for {
			select {
			case temp := <-n:
				fmt.Println("receiving from channel", temp)
			case <-stop1:
				// do cleanup
				time.Sleep(100 * time.Millisecond)
				fmt.Println("performed cleanup")
				stop1 <- true
				return
			}
		}
	}()
}
