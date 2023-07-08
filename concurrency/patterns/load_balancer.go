package main

import (
	"container/heap"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Request struct { // sends request to balancer
	fn func() int
	c  chan int
	id int
}

type Worker struct {
	requests chan Request // buffered channel
	pending  int          // count of pending tasks
	index    int          // index in heap
}

type Pool []*Worker

type Balancer struct { // sends request to worker
	pool Pool
	done chan *Worker // report work done by worker
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Push(x interface{}) {
	p = append(p, x.(*Worker))
}
func (p Pool) Pop() interface{} {
	lastIndex := len(p) - 1
	lastElement := p[lastIndex]
	p = p[:lastIndex]
	return lastElement
}

// load generator
func requester(work chan<- Request, done chan bool, stopSignal chan bool) {
	c := make(chan int)
	start := time.Now()

	for {
		select {
		case <-done:
			fmt.Println("Interrupt Signal Received, send signal to balancer")
			stopSignal <- true
			return
		case work <- Request{c: c,
			fn: workFunc}:
			fmt.Println("Request Arrived")
			result := <-c
			furtherProcessResult(result, start)
		}
	}
}

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

func workFunc() int {
	return generateInteger()
}

func furtherProcessResult(output int, start time.Time) {
	fmt.Println(fmt.Sprintf("Request %d processed in %v ", output, time.Since(start)))
}

func (b *Balancer) dispatch(req Request) {
	// get worker
	worker := heap.Pop(&b.pool).(*Worker)
	// send it request
	worker.requests <- req
	worker.pending++
	heap.Push(&b.pool, worker)
}
func (b *Balancer) completed(w *Worker) {
	for i, worker := range b.pool {
		if worker.index == w.index {
			w.pending--
			heap.Fix(&b.pool, i)
		}
	}
}

func (b *Balancer) balance(work chan Request, stopSignal <-chan bool) {
	for {
		select {
		case req := <-work: // received a request
			fmt.Println("Received request")
			go b.dispatch(req) // so send it to a worker
		case w := <-b.done: // worker has finished
			go b.completed(w) // so update info
		case <-stopSignal:
			return
		}
	}
}

func (w *Worker) done(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		fmt.Println(fmt.Sprintf("Worker %d completed task", w.index))
		time.Sleep(2 * time.Millisecond)
		done <- w
	}
}

func main() {

	// create channel to accumulate work
	work := make(chan Request)

	// create workers
	var pool Pool
	workers := 10
	for i := 0; i <= workers; i++ {
		pool = append(pool, &Worker{
			requests: make(chan Request),
			index:    i,
			pending:  0,
		})
	}
	heap.Init(pool)

	// result chan
	doneCh := make(chan *Worker)
	balancer := Balancer{
		pool: pool,
		done: doneCh,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	// checks if it receives interrupt signal
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// stopSignal for requester
	stopSignalP := make(chan bool)

	// stopSignal for balancer
	stopSignal := make(chan bool)

	go requester(work, stopSignalP, stopSignal)

	go balancer.balance(work, stopSignal)

	for _, worker := range balancer.pool {
		go func(workerr *Worker) {
			workerr.done(doneCh)
		}(worker)
	}

	<-done
	// send signal to producer to stop
	stopSignalP <- true
	fmt.Println("All done and dusted")

}
