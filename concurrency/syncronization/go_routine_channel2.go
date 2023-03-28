package main

import (
	"fmt"
)

// sender receiver both are go routines
// sender sends data async, receiver receives data async
// before sync main func exits, sender and receiver are not done
func main() {
	ch := make(chan bool)

	doWork := func(ch chan bool) {
		fmt.Println("started")
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}

		ch <- true
	}

	go doWork(ch)

	// not a blocking call so system exits before receiver is done
	go func() {
		fmt.Println(<-ch)
	}()

	// time.Sleep(time.Second)

}
