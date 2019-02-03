package pool

import (
	"fmt"
	"sync"
)

// Work - data to process
type Work struct {
	Before string
	After  string
}

// Worker - process Work from an input channel and places the results to
// and output channel.  Gracefully stops processing when stop channel has
// something on it.
type Worker struct {
	waiter *sync.WaitGroup
	// read from channel
	todo <-chan Work

	// write to channel
	results chan<- Work

	// stop processing and terminate the worker
	stop <-chan struct{}
}

// Start - start processing the input channel
func (w *Worker) Start() {
	var (
		work Work
		stop bool
	)
	// When the worker exits, decrement the waitgroup
	defer w.waiter.Done()

	for {
		if stop {
			// break out of the for loop
			// ending the worker
			break
		}

		// check for data on either channel
		select {
		case work = <-w.todo:
			fmt.Printf("TODO: Goto work: %s\n", work.Before)
		case <-w.stop:
			fmt.Println("worker told to stop")
			stop = true
		}
	}
}

// Pool - a group of workers
type Pool struct {
	waiter  *sync.WaitGroup
	workers []*Worker
	Todo    chan<- Work
	Results chan Work

	// sending to this channel will gracefully stop the workers
	Stop chan struct{}

	// Workers - number of Workers to start
	Workers int
}

// New - creates and starts workers to process the todo channel
func New(workers int) *Pool {
	wg := &sync.WaitGroup{}
	input := make(chan Work)
	output := make(chan Work)
	stop := make(chan struct{})
	pool := &Pool{waiter: wg,
		Todo:    input,
		Results: output,
		Stop:    stop,
	}

	for i := 0; i <= workers; i++ {
		w := &Worker{
			waiter:  wg,
			todo:    input,
			results: output,
			stop:    stop,
		}
		pool.workers = append(pool.workers, w)

		// Increment the wait group
		pool.waiter.Add(1)

		// start the worker in it's own go routine
		go pool.workers[i].Start()
	}
	return pool
}

// Halt - send an empty struct{} to each worker
// to gracefully shutdown the pool.
func (p *Pool) Halt() {
	var halt struct{}
	for range p.workers {
		p.Stop <- halt
	}
	p.waiter.Wait()
	close(p.Results)
}
