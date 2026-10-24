package main

import "context"

type Output struct {
	value any
	error error
}

/*
 * NOTE: Returns immediately with a function that takes in a receive only channel
 * so task:=async(fn)
 * returns the closure and not the inner body
 */
func async(ctx context.Context, fn func(ctx context.Context) Output) func() <-chan Output {
	return func() <-chan Output {
		/*
			 * NOTE: this runs when the function is called
				* so task()
		*/
		out := make(chan Output, 1)
		go func() {
			out <- fn(ctx)
		}()
		return out
	}
}

func await(in <-chan Output) Output {
	return <-in
}
