package concurrency_problems

import (
	"testing"
)

var done1 = make(chan bool)
var msgs1 = make(chan int)

func produce1() {
	for i := 0; i < 10; i++ {
		msgs1 <- i // send message
	}
	close(msgs1)
	done1 <- true //send signal once producer is done
}

func consume1() {
	for {
		// receiving from a closed channel returns the zero value immediately
		// it keeps on running infinitely
		msg := <-msgs1 // receive message
		println(msg)
	}
}

func validateFlow1() {
	go produce1()
	go consume1()
	<-done1 // wait for done signal
}

func TestProducerFlow1(t *testing.T) {
	validateFlow1()
}
