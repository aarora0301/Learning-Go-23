package concurrency_problems

import "fmt"

func worker(id int, tasks <-chan int, done chan<- bool) {
	for task := range tasks {
		fmt.Printf("worker %d started task %d", id, task)
		// do something
		done <- true
	}
}

func setupTasks() {
	maxConcurrency := 3
	totalTasks := 10
	tasks := make(chan int)
	done := make(chan bool)

	// start workers
	for i := 1; i <= maxConcurrency; i++ {
		go worker(i, tasks, done)
	}

	// submit tasks
	go func() {
		for i := 1; i <= totalTasks; i++ {
			tasks <- i
		}
		// close channel after tasks are sent
		close(tasks)
	}()

	for i := 1; i <= totalTasks; i++ {
		<-done // wait for all tasks to complete
	}
}
