package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		fmt.Println("Produced:", num)
		ch <- num
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch chan int) {
	for item := range ch {
		fmt.Println("Consumed:", item)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
