package main

import (
	"fmt"

	"github.com/kirk-patton/example-go/channels/workerpool/pool"
)

func main() {
	workerpool := pool.New(5)

	// Place some work on the input channel
	for i := 1; i <= 100; i++ {
		todo := pool.Work{Before: "hello"}
		workerpool.Todo <- todo
	}
	workerpool.Halt()

	for result := range workerpool.Results {
		fmt.Printf("Done: %s", result.Before)
	}

}
