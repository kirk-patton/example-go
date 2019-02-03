package main

import (
	"fmt"
	"time"

	"github.com/kirk-patton/example-go/channels/workerpool/pool"
)

func main() {
	// start the workerpool so there is something pprocessing the input channel
	workerpool := pool.New(5)

	// we need to have something listening to the workerpool.results
	// channel before the worker goroutne tries to write to the channel
	// otherwise a deadlock will occure
	go func() {
		for result := range workerpool.Results {
			fmt.Printf("Done: %s", result.Before)
		}
	}()

	// Place some work on the input channel
	for i := 1; i <= 100; i++ {
		todo := pool.Work{Before: "hello"}
		workerpool.Todo <- todo
	}

	go func() {
		time.Sleep(5)
		workerpool.Halt()
	}()

	for result := range workerpool.Results {
		fmt.Printf("Done: %s\n", result.Before)
	}

}
