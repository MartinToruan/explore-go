package main

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoRoutines int = 25
	maxPool       int = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	fmt.Println("Close: Connection ", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Println("Create New Connection: ", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoRoutines)

	pool, err := New(createConnection, uint(maxPool))
	if err != nil {
		fmt.Println(err)
		return
	}

	for queryID := 0; queryID < maxGoRoutines; queryID++ {
		go func() {
			performQuery(queryID, pool)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("All queries have been executed...")
	pool.Close()
	fmt.Println("Program shutdown.")
}

func performQuery(queryID int, pool *Pool) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	conn, err := pool.Acquire()
	if err != nil {
		fmt.Println("error while do query: ", err)
	}
	defer pool.Release(conn)

	fmt.Printf("QueryID : %d, ConnectionID: %d\n", queryID, conn.(*dbConnection).ID)
}
