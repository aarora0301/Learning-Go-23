package main

import (
	"fmt"
	"time"
)

func main() {

	numJobs := 5
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int)

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

	for job := 1; job <= numJobs; job++ {
		jobs <- job
	}

	close(jobs)

	//this fails because it works like an infinite loop
	// tries to receive from channel infinitely
	//for elem := range results {
	//	fmt.Println("result: ", elem)
	//}

	for i := 0; i < numJobs; i++ {
		fmt.Println("result: ", <-results)
	}
}
