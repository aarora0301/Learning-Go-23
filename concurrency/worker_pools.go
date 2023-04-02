package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int)
	results := make(chan int)

	numJobs := 5
	numWorkers := 3

	workers := func(id int, jobs <-chan int, results chan<- int) {
		for job := range jobs {
			fmt.Println("worker", id, "started  job", job)
			time.Sleep(time.Second) // simulate work
			fmt.Println("worker", id, "finished job", job)
			results <- job * 2
		}
	}

	// create a pool of 3 workers
	for i := 1; i <= numWorkers; i++ {
		go workers(i, jobs, results)
	}

	//deadlock ,
	// noOfSender:=5
	// receiver:=3
	// so after three packets all recievers are busy
	// so sender is not able to insert in channel
	// so it makes sense in case of worker pools we use buffered channels
	// works get accumulated in buffered channel and then any receiver can pick from it once
	// it is free

	for job := 1; job <= numJobs; job++ {
		jobs <- job
	}

	close(jobs)

	for i := 0; i < numJobs; i++ {
		fmt.Println("result: ", <-results)
	}

}
