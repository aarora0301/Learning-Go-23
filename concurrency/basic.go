package main

import "log"

// go routine exits when main function exits
func main() {

	text := func(text string) {
		log.Println("started :", text)
		for i := 0; i < 3; i++ {
			log.Println(text, i)
		}
	}

	text("Main")

	go text("Goroutine 1")

	go text("Goroutine 2")

}
