package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {

		// uncomment above code and use go vet to evaluate the code

		//wg.Add(1)
		//
		//
		//go func() {
		//	defer wg.Done()
		//	fmt.Println("Hello1", i)
		//}()

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Hello2", i)
		}(i)

	}
	wg.Wait()

}
