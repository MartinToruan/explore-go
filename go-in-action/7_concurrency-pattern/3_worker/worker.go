package main

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxWorker int) *Pool {
	pool := Pool{
		work: make(chan Worker),
	}

	pool.wg.Add(maxWorker)

	for i := 0; i < maxWorker; i++ {
		go func() {
			for worker := range pool.work {
				worker.Task()
			}
			pool.wg.Done()
		}()
	}

	return &pool
}

func (p *Pool) Run(task Worker) {
	p.work <- task
}

func (p *Pool) Shutdown() {
	close(p.work)

	p.wg.Wait()
}
