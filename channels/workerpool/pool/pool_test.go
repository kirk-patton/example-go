// Example of internal unit test
// all public & private variables / types / functions etc... visable to test logic
package pool

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	wg := &sync.WaitGroup{}
	todo := make(chan Work)
	result := make(chan Work)
	stop := make(chan struct{})
	w := &Worker{
		waiter:  wg,
		todo:    todo,
		results: result,
		stop:    stop,
	}

	// start the worker
	go w.Start()

	// create some work
	expect := "HELLO"
	x := Work{Before: strings.ToLower(expect)}

	// place it on the channel
	todo <- x

	// close the input channel
	// the worker is coded to shut down when the input channel closes
	close(todo)

	// read back the result
	r := <-result

	msg := fmt.Sprintf("Expected: %s, Got: %s", expect, r.After)
	if r.After != expect {
		t.Error(msg)
	}

	t.Log(msg)
	wg.Wait()
}
