package main

import (
	"fmt"
	"time"
)

func printNumbers(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go printNumbers(ch)

	for num := range ch {
		fmt.Println(num)
	}
}
