package main

import "context"

type PromiseAllOutput struct {
	values []any
	error  error
}

type indexedOutput struct {
	index  int
	output Output
}

func promiseAll(cancel context.CancelFunc, promises ...<-chan Output) <-chan PromiseAllOutput {
	output := make(chan PromiseAllOutput, 1)
	go func() {
		results := make([]any, len(promises))
		done := make(chan indexedOutput, len(promises))
		for i, promise := range promises {
			go func(index int, promise <-chan Output) {
				done <- indexedOutput{
					index:  index,
					output: <-promise,
				}
			}(i, promise)
		}
		for range promises {
			result := <-done
			if result.output.error != nil {
				cancel()
				output <- PromiseAllOutput{
					error: result.output.error,
				}
				return
			}
			results[result.index] = result.output.value
		}
		output <- PromiseAllOutput{
			values: results,
		}
	}()
	return output
}

func awaitPromiseAll(in <-chan PromiseAllOutput) PromiseAllOutput {
	return <-in
}
