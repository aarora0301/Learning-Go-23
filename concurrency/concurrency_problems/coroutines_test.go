package concurrency_problems

import (
	"fmt"
	"testing"
)

func integers() chan int {
	result := make(chan int)
	count := 0

	// called once will always run in background
	go func() {
		for {
			count++
			result <- count
		}
	}()
	return result
}

var y = integers()

func generateInteger() int {
	return <-y
}

func TestCoroutines(t *testing.T) {
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}
