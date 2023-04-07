package main

import (
	"fmt"
	"runtime"
)

// no. of threads in machine
func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())
}
