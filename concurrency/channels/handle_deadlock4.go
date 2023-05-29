package main

import (
	"fmt"
	"sync"
)

func generateIntegers2(n int) chan int {
	result := make(chan int)
	var wg sync.WaitGroup

	wg.Add(n) // Add the number of goroutines to wait for

	for i := 0; i < n; i++ {
		go func(num int) {
			result <- num
			wg.Done() // Notify the wait group that this goroutine has completed
		}(i)
	}

	go func() {
		wg.Wait()     // Wait for all goroutines to complete
		close(result) // Close the result channel after all values have been sent
	}()

	return result
}

func main() {
	for num := range generateIntegers2(5) {
		fmt.Println(num)
	}
}
