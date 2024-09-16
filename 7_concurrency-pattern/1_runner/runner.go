package main

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"time"
)

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	task []func(int)
}

func New(timeout int) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(time.Duration(timeout) * time.Second),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.task = append(r.task, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.task {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
