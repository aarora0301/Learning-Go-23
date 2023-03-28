package main

import (
	"fmt"
	"io"
	"os"
)

func echo(reader io.Reader, writer io.Writer) {
	fmt.Println("inside")
	io.Copy(writer, reader)
}

func main() {
	go echo(os.Stdin, os.Stdout)
	//time.Sleep(30 * time.Second) // go routine reads from args for 30 seconds, run until main/caller function runs
	//fmt.Println("time over")

}
