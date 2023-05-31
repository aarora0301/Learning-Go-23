package problems

import (
	"fmt"
	"sync"
	"testing"
)

// The code for “two threads printing alternate parity numbers”
func even_odd(N int) {
	ch := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2 * N)

	go func() {
		for i := 0; i < N; i++ {
			defer wg.Done()
			ch <- struct{}{}
			if i%2 != 0 {
				fmt.Println("odd", i)
			}
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			defer wg.Done()
			<-ch // blocks until other go routine is processed

			if i%2 == 0 {
				fmt.Println("even", i)
			}

		}
	}()

	wg.Wait()
}

func TestFlow(t *testing.T) {
	even_odd(15)
}
