package main

import (
	"algorithm-coding/concurrency"
	"time"
)

func main() {
	ch := make(chan int, 64)

	go concurrency.Producer(3, ch)
	go concurrency.Producer(5, ch)
	go concurrency.Consumer(ch)

	time.Sleep(5 * time.Second)

}
