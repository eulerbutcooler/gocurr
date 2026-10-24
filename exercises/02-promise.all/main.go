package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	start := time.Now()
	p1 := async(ctx, func(ctx context.Context) Output {
		time.Sleep(300 * time.Millisecond)
		return Output{value: "p1"}
	})()

	p2 := async(ctx, func(ctx context.Context) Output {
		time.Sleep(100 * time.Millisecond)
		return Output{value: "p2"}
	})()

	p3 := async(ctx, func(ctx context.Context) Output {
		time.Sleep(200 * time.Millisecond)
		return Output{value: "p3"}
	})()

	result := awaitPromiseAll(promiseAll(cancel, p1, p2, p3))
	fmt.Println("elapsed: ", time.Since(start))
	if result.error != nil {
		fmt.Println("error: ", result.error)
		return
	}
	fmt.Println("results: ", result.values)
}
