package concurrency_problems

import (
	"fmt"
	"sync"
	"testing"
)

type WaitGroup struct {
	counter int
	wait    bool
	done    chan bool
	mu      sync.Mutex
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		done: make(chan bool),
	}
}

func (wg *WaitGroup) Add(n int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter += n

	if wg.counter < 0 {
		panic("negative wait counter")
	}

	if wg.counter == 0 && wg.wait {
		wg.done <- true
		wg.wait = false
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.mu.Lock()

	if wg.counter == 0 {
		wg.mu.Unlock()
		return
	}
	wg.wait = true
	wg.mu.Unlock()
	<-wg.done
}

func TestWaitGroupFlow(t *testing.T) {
	t.Run("validate wait group flow", func(t *testing.T) {
		wg := NewWaitGroup()

		for i := 0; i < 10; i++ {
			go func(n int) {
				wg.Add(1)
				fmt.Println(n)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}
