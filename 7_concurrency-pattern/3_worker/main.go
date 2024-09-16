package main

import (
	"fmt"
	"sync"
	"time"
)

type namePrinter struct {
	name string
}

func (np *namePrinter) Task() {
	fmt.Println(np.name)
	time.Sleep(time.Duration(500) * time.Millisecond)
}

var names = []string{"Budi", "Rudi", "Sudi", "Midi", "Kidi"}

func main() {
	var wg sync.WaitGroup
	wPool := New(8)

	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, n := range names {
			go func() {
				wPool.Run(&namePrinter{name: n})
				wg.Done()
			}()
		}
	}

	wg.Wait()
	wPool.Shutdown()
	fmt.Println("Program is finished")
}
