package main

import (
	"fmt"
)

func fibonacciSeries(n int) []int {
	if n <= 0 {
		return []int{}
	}
	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}
func main() {
	fmt.Println(fibonacciSeries(6)) // [0 1 1 2 3 5]

}
