package main

import (
	"fmt"
	"sync"
)

func generateIntegers1(n int) chan int {
	result := make(chan int)
	var wg sync.WaitGroup

	wg.Add(n) // Add the number of goroutines to wait for

	for i := 0; i < n; i++ {
		go func(num int) {
			result <- num
			wg.Done() // Notify the wait group that this goroutine has completed
		}(i)
	}

	// deadlock here as main is indefinitely waiting
	//for all the goroutines to complete
	// including the one that should close the channel
	wg.Wait()
	close(result)

	return result
}

func main() {
	for num := range generateIntegers1(5) {
		fmt.Println(num)
	}
}
