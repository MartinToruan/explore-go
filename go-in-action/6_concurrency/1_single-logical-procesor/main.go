package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Using 1 Logical Processor
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")

}
