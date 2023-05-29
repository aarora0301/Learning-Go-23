package concurrency_problems

import (
	"fmt"
	"testing"
)

//normal syncronous Func

func getSquare(num int) int {
	return num * num
}

func product(num1 int, num2 int) int {
	return num1 * num2
}

// Improving this function using go routines
// such that parallely compute num1 and num2
// and then compute the product
func SquaredProduct(a int, b int) int {
	num1 := getSquare(a)
	num2 := getSquare(b)
	result := product(num1, num2)
	fmt.Println(result)
	return result
}

type future chan int

// consume from channel
// produce output from channel
func GetSquareAsync(a future) future {
	result := make(future)
	go func() {
		result <- getSquare(<-a)
	}()
	return result
}

func GetSquare(a int) int {
	return <-GetSquareAsync(Promise(a))
}

// add input to channel
func Promise(a int) future {
	// it has to be a buffered channel otherwise it will block
	result := make(future, 1)
	result <- a
	return result
}

func ProductAsync(a future, b future) future {
	result := make(future)
	go func() {
		result <- product(<-a, <-b)
	}()
	return result
}

func Product(a int, b int) int {
	return <-ProductAsync(Promise(a), Promise(b))
}

// runs a and b parallely
// and then computes the product
func SquaredProductOutput(a, b int) {
	a1 := GetSquare(a)
	b1 := GetSquare(b)
	fmt.Println(Product(a1, b1))
}

func TestSquares(t *testing.T) {
	fmt.Println("Normal Syncronous Function")
	SquaredProduct(3, 4)
	fmt.Println("Asyncronous Function")
	SquaredProductOutput(3, 4)
}
