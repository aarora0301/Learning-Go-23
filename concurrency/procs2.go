package main

import "fmt"

func main() {

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

// Run it as

// 1. GOMAXPROCS=1 go run main.go

//2. GOMAXPROCS=2 go run main.go

// In the first run, at most one goroutine was executed at a time. Initially, it was the main goroutine,
//which prints ones (output trimmed using ...). After a period of time, the Go scheduler put it to sleep
//and woke up the goroutine that prints zeros, giving it a turn to run on the OS thread.
//
//In the second run, there were two OS threads available, so both goroutines were trying to utilise the thread
//simultaneously, printing digits at some interval. /In yo
