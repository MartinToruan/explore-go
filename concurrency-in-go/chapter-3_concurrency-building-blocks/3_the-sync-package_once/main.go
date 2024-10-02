package main

import (
	"fmt"
	"sync"
)

func main() {
	// Example1()
	Example2()
}

/*
You can see in this example that the increment function is only invoked once.
Resulting the count value is still 1 (not a 100)
*/
func Example1() {
	var count int

	var increment = func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	for i := 0; i < 100; i++ {
		increments.Add(1)
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

/*
In this example shows that once package only counts the number of Do functions is called, not how many unique functions passed into Do are called.
So, this example only invoked the increment function, and will pass the decrement function invocation.
*/
func Example2() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count: %d\n", count)
}
