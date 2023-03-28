package main

import "fmt"

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

	// blocking call until data is received
	fmt.Println(<-ch)

}
