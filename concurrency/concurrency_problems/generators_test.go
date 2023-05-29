package concurrency_problems

import (
	"fmt"
	"testing"
)

func generateIntegers(n int) chan int {
	result := make(chan int)

	// this is analogous to wait group
	sem := make(semaphore)

	for i := 0; i < n; i++ {
		go func(num int) {
			result <- num
			sem.Signal() // just like wg.Done()
		}(i)
	}

	go func() {
		sem.Wait(n) // just like wg.Wait()
		close(result)
	}()

	return result
}

func TestGenerators(t *testing.T) {

	//  here main waits indefinitely for the result channel
	// to be closed, which never happens.
	for num := range generateIntegers(5) {
		fmt.Println(num)
	}
}
