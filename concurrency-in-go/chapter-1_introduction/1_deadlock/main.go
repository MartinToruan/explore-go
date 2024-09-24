package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

/*
This is deadlock because
  - goroutine 1 locks the valOne, and goroutine 2 locks the valTwo
  - After 2 seconds, goroutine 1 tries to lock valTwo (which is currenctly locked by goroutine 2)
*/
var printSum = func(v1, v2 *value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(2 * time.Second)

	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Printf("sum=%v\n", v1.value+v2.value)
}

func main() {
	var valOne, valTwo value
	wg.Add(2)
	go printSum(&valOne, &valTwo)
	go printSum(&valTwo, &valOne)

	wg.Wait()
}
