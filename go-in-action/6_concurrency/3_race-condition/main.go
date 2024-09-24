package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	// The result will be 2
	// Expected result is 4
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// Ask the go runtime scheduler to held the thread and let other goroutines to run,
		// thus the current gr will be placed back in queue
		runtime.Gosched()

		value++

		counter = value
	}
}
