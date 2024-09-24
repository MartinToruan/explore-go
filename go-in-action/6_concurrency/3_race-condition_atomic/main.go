package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	// The result will be 4
	// Expected result is 4
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Atomic AddInt64 will make sure that only 1 Go routine will be able to increment the value at a time.
		// If there's 2 Go Routines or more trying to increment the value, it will synchronize automatically
		atomic.AddInt64(&counter, 1)

		// Ask the go runtime scheduler to held the thread and let other goroutines to run,
		// thus the current gr will be placed back in queue
		runtime.Gosched()
	}
}
