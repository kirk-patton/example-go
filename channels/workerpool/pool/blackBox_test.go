package pool_test

import (
	"fmt"

	"github.com/kirk-patton/example-go/channels/workerpool/pool"
)

// ExampleNew - is a unit test that will show up as an example in our generated godoc
// The required keyword, "Example", must be followed by the name of an existing function in
// the package under test
func ExampleNew() {
	count := 10
	workerpool := pool.New(count)

	// Place some work on the input channel
	for i := 1; i <= count; i++ {
		todo := pool.Work{Before: "hello"}
		workerpool.Todo <- todo
	}
	go workerpool.Halt()

	for result := range workerpool.Results {
		fmt.Printf("Before: %s => After: %s\n", result.Before, result.After)
	}
	// Output:
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
	// Before: hello => After: HELLO
}
