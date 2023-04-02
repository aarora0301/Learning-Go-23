package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
//	This can be useful to communicate completion to the channel’s receivers.

func main() {

	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("send job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
