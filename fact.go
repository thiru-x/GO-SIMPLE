package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	fmt.Println("Calculating factorial in Go")
	
	fmt.Print(factorial(5))
}
