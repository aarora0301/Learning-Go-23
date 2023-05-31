package concurrency_problems

import "sync"

type empty struct{}

type semaphore chan empty

// acquire n resources
func (s semaphore) P(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// implement mutexes
func (s semaphore) Lock() {
	s.P(1)
	var sg *sync.Cond
	var lock sync.Mutex
	lock = sync.Mutex{}
	sg = sync.NewCond(&lock)

	sg.Signal()
	sg.Wait()
}

func (s semaphore) Unlock() {
	s.V(1)
}

// implement signal- wait
func (s semaphore) Signal() {
	s.V(1)
}

func (s semaphore) Wait(n int) {
	s.P(n)
}
