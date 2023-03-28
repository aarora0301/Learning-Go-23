package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

func print(reader io.Reader, writer io.Writer) {
	fmt.Println("inside")

}

func main() {
	go print(os.Stdin, os.Stdout)
	runtime.Gosched() // go scheduler  may then run one or more of them (other go routines) before coming back to
	// this function, alternative to time.Sleep()
	//
}
