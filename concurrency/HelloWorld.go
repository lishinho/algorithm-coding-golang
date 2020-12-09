package concurrency

import (
	"fmt"
	"sync"
)

// use mutex to do helloWorld
func MuHelloWorld() {
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		mu.Lock()
		go func() {
			fmt.Printf("Hello, world in mutex %d \n", i)
			mu.Unlock()
		}()
	}

	mu.Lock()
}

func ChHelloWorld() {
	done := make(chan int, 10) // channel with buffer

	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Printf("Hello world in channel %d \n", i)
			done <- 1
		}()
	}

	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func WgHelloWorld() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Printf("Hello world in wait group %d \n", i)
			wg.Done()
		}()
	}

	wg.Wait()
}
