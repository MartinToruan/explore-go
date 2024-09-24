package main

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("pool has been closed")

type Pool struct {
	mu       sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value is too small")
	}

	return &Pool{
		resource: make(chan io.Closer, size),
		factory:  fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case res, ok := <-p.resource:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquire Shared Resource")
		return res, nil
	default:
		fmt.Println("Creating a new resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		_ = r.Close()
		return
	}

	select {
	case p.resource <- r:
		fmt.Println("Release: In Queue")
	default:
		_ = r.Close()
		fmt.Println("Release: Closing")
	}
}

func (p *Pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}

	// Set the pool to the close
	p.closed = true

	// close the channel so no goroutine can acquire a new resource
	close(p.resource)

	// Flush the remaining resource
	for r := range p.resource {
		_ = r.Close()
	}
}
