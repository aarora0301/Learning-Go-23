package main

import (
	"fmt"
	"time"
)

//Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a
//channel that will be notified at that time.

// If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is
//that you can cancel the timer before it fires

func main() {

	start := time.Now()
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C

	fmt.Println("Timer 1 fired") // executes after 2 seconds

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop() // stops timer before a second has passed

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	fmt.Println("Time elapsed: ", time.Since(start))
}
