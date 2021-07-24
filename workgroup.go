/*
Package workgroup provides functions to manage the lifetime of a related set of goroutines.
*/
package workgroup

import (
	"context"
	"sync"
)

/*
Run will execute all functions in the slice passed to it in individual goroutines.
Run blocks until all the goroutines spawned have ended execution.
The first function to return a non-nil value will trigger the end of execution of all other goroutines spawned.
The return value from the first function will be returned to the caller of Run.
If all none of the functions return a non-nil value, all the spawned goroutines will naturally end execution and
Run will return nil.
*/
func Run(fns []func() interface{}) interface{} {
	// Make channel to receive output from the goroutines.
	// We're using a channel instead of setting a variable to ensure that the first value returned by a goroutine
	// is what this function returns.
	output := make(chan interface{}, 1)
	defer close(output)

	// Make a wait group and set delta to number of functions to run.
	var wg sync.WaitGroup
	wg.Add(len(fns))

	// Make a context with cancel to stop all other goroutines once a function ends it's execution.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run all functions in goroutines
	for _, fn := range fns {
		go func(fn func() interface{}, output chan<- interface{}) {
			// Reduce delta of wait group by 1 when execution of goroutine is done or cancelled.
			defer wg.Done()

			select {
			// End execution when cancel is called by another goroutine.
			case <-ctx.Done():
				return

			// If cancel has not been called, run the function.
			default:
				op := fn()
				// If returned value is not nil, send output to channel.
				if op != nil {
					output <- op
					cancel()
				}
			}
		}(fn, output)
	}

	// Wait for all goroutines to end execution.
	wg.Wait()

	// Return output from first function to end execution.
	// If no output has been sent to channel, return nil.
	if len(output) > 0 {
		return <-output
	} else {
		return nil
	}
}
