package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(numberGoroutines)

	chLoad := make(chan string, taskLoad)

	// Spawn worker
	for i := 1; i <= 4; i++ {
		go worker(chLoad, i)
	}

	// Spawn Task
	for i := 1; i <= 10; i++ {
		chLoad <- fmt.Sprintf("Task : %d", i)
	}

	close(chLoad)

	wg.Wait()

}

func worker(task chan string, worker int) {
	defer wg.Done()

	for {
		t, ok := <-task
		if !ok {
			fmt.Println("Shutting Down")
			return
		}
		fmt.Printf("Worker %d : Started %s\n", worker, t)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker %d : Completed %s\n", worker, t)
	}
}
