package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	runOneTimeBroadcast()
}

type Button struct {
	Clicked *sync.Cond
}

var subscribeOneTime = func(c *sync.Cond, fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)

	go func() {
		goroutineRunning.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fn()
	}()

	goroutineRunning.Wait()
}

/*
This function simulate how to use sync.Cond to send a broadcast one time.
As you can see, Online 25 we call w.Wait(). This command instruct the goroutine to wait for a broadcast message from the main routine.
On line 57, you can see we call the Broadcast that makes all the goroutine do it's task [fn()]
*/
func runOneTimeBroadcast() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribeOneTime(button.Clicked, func() {
		fmt.Println("Maximizing Window.")
		clickRegistered.Done()
	})
	subscribeMultipleTimes(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box.")
		clickRegistered.Done()
	})

	subscribeMultipleTimes(button.Clicked, func() {
		fmt.Println("Mouse Clicked.")
		clickRegistered.Done()
	})

	time.Sleep(5 * time.Second)
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}

var subscribeMultipleTimes = func(c *sync.Cond, fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)

	go func() {
		goroutineRunning.Done()
		for {
			c.L.Lock()
			c.Wait()
			fn()
			c.L.Unlock()
		}
	}()

	goroutineRunning.Wait()
}

func runMultipleTimeBroadcast() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(6)

	subscribeMultipleTimes(button.Clicked, func() {
		fmt.Println("Maximizing Window.")
		clickRegistered.Done()
	})

	subscribeMultipleTimes(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box.")
		clickRegistered.Done()
	})

	subscribeMultipleTimes(button.Clicked, func() {
		fmt.Println("Mouse Clicked.")
		clickRegistered.Done()
	})

	time.Sleep(5 * time.Second)
	button.Clicked.Broadcast()
	time.Sleep(5 * time.Second)
	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
