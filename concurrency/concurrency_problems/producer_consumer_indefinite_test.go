package concurrency_problems

import (
	"testing"
)

var done = make(chan bool)
var msgs = make(chan int)

func produce() {
	for i := 0; i < 10; i++ {
		msgs <- i // send message
	}
	close(msgs)
	//done <- true // until done channel receives something, main will not exit as line 30 is blocking
}

func consume() {
	for {
		// receiving from a closed channel returns the zero value immediately
		// it keeps on running infinitely
		msg := <-msgs // receive message
		println(msg)
	}
}

func validateFlow() {
	go produce()
	go consume()
	<-done // wait for done signal
}

func TestProducerFlow(t *testing.T) {
	validateFlow()
}
