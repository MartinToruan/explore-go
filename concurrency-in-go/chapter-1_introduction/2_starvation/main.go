package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sharedLock sync.Mutex

const rt = 1 * time.Second

var greedyWorker = func() {
	defer wg.Done()

	var count int
	for begin := time.Now(); time.Since(begin) <= rt; {
		sharedLock.Lock()
		time.Sleep(3 * time.Nanosecond)
		sharedLock.Unlock()
		count++
	}
	fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
}

var politeWorker = func() {
	defer wg.Done()

	var count int
	for begin := time.Now(); time.Since(begin) <= rt; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()
		count++
	}
	fmt.Printf("Polite worker was able to execute %v work loops\n", count)
}

// This sample code shows that greedy worker will get more CPU time (align with number of work is greater than polite worker)
// This is because he locks the CPU longer
// So, be careful to lock the process, because it will block the other. Make sure you only lock your critical part of the code.

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
