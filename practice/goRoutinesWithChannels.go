package main

import (
	"fmt"
	"sync"
	"time"
)

func GoRoutinesWithChannels() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// Start goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	// Send values to the channel
	for i := 1; i <= 25; i++ {
		ch <- i
	}

	// Close the channel to signal no more values
	close(ch)

	// Wait for all goroutines to finish
	wg.Wait()
}

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	j := 1
	for v := range ch {
		time.Sleep(time.Millisecond)
		fmt.Printf("%d.%d got %d\n", id, j, v)
		j++
	}
}