package main

import "fmt"

func generateIntegers(n int) chan int {
	result := make(chan int)

	for i := 0; i < n; i++ {
		go func(num int) {
			result <- num
		}(i)
	}

	return result
}

func main() {

	//  here main waits indefinitely for the result channel
	// to be closed, which never happens.

	// A channel should be closed by the goroutine that writes to it.
	for num := range generateIntegers(5) {
		fmt.Println(num)
	}
}
