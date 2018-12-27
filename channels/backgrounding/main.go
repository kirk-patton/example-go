package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Data - some data to process
type Data struct {
	Before string
	After  string
}

func main() {
	// init our random number generator
	rand.Seed(time.Now().Unix())

	// blocking channels - no buffer specified
	// only a single Data instance can be placed on the channel
	// at a time
	input := make(chan Data)
	output := make(chan Data)

	// this runs seprately from the rest of the instructions in main
	// it blocks until main() sends it something to do. It is
	// started first so that something is listening on the input
	// channel
	go func(in <-chan Data, out chan<- Data) {
		// close the output channel when this
		// function exits. Notice that we
		// specified that it is read from with "<-chan"
		// and that out is written to with "chan<-".
		// the arrow always points to the left.
		defer close(out)

		// iterate over the input channel
		// changing the case of the Before field
		// and place the modified struct on the output
		// channel
		for item := range input {
			// Note: this loop terminate when the input
			// channel is closed
			item.After = strings.ToUpper(item.Before)

			// place the updated item on the output channel
			out <- item
		}

	}(input, output)

	// Now, lets give the first go routine something to process by starting a
	// second go routine that generates lowercase letters and places them in our Data
	// struct.  This tool runs in the background...
	go func(in chan<- Data) {
		// loop forever
		for {
			// initialize x as a "Data" type with the "Brefore" field set to a random lowercase letter
			x := Data{Before: randChar()}

			// place it on the channel
			in <- x
		}
	}(input)

	// Process the output channel
	for result := range output {
		fmt.Printf("Before: %s => After: %s\n", result.Before, result.After)
	}

}

// return a random lowercase ascii character
func randChar() string {
	// ascii 97-122 lowerdcase a-z
	c := rand.Intn(122-97) + 97 + 1
	return fmt.Sprintf("%s", string(c))
}
