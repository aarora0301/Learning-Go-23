package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// run file as go run -race atomic_counter.go
func main() {

	var ops uint64
	var counter uint64
	var wg sync.WaitGroup

	// using atomic counter
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			defer wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("ops:", ops)

	// using non atimoic counter
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				counter++
			}
			defer wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("counter:", counter)

}
