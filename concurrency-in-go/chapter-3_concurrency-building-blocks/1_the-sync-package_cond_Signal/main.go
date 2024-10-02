package main

import (
	"fmt"
	"sync"
	"time"
)

/*
In this example, you can see that our for loop is waiting the Signal from the removeQueue go routine
When it's got the signal, it will continue the work
*/
func main() {
	c := sync.NewCond(&sync.Mutex{})

	queue := make([]interface{}, 0, 10)

	removeQueue := func(delay time.Duration) {
		time.Sleep(delay)

		// Critical Section is started here
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		// Another Critical Part
		c.L.Lock()
		for len(queue) == 2 {
			fmt.Println("\tWaiting")
			c.Wait()
			fmt.Println("\tDone Waiting")
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeQueue(5 * time.Second)
		c.L.Unlock()
	}
}
