package main

import (
	"fmt"
)

func main() {

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("Unknown type", t)
		}

	}

	whatAmI(1 == 1)
	whatAmI(1)

}
