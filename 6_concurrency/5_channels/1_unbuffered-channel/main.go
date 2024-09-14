package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// Add a count of two, one for each go routines
	wg.Add(2)

	// Launch 2 players
	go player("Yuda", ch)
	go player("MC", ch)

	// Start the set
	ch <- 1

	// wait for the game to finish
	wg.Wait()

}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court

		if !ok {
			// If the channel is closed, the player won
			fmt.Printf("Player %s won\n", name)
			return
		}

		// Pick a random number
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)

		ball++

		court <- ball
	}
}
