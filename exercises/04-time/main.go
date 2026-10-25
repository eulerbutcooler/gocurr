package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func throttle(n int, fn func() int) (handle func() error, wait func()) {
	sema := make(chan struct{}, n)
	handle = func() error {
		select {
		case sema <- struct{}{}:
			go func() {
				fn()
				<-sema
			}()
			return nil
		default:
			return errors.New("BUSY")
		}
	}
	wait = func() {
		for range n {
			sema <- struct{}{}
		}
	}
	return handle, wait
}

func work() int {
	if rand.Intn(10) < 8 {
		time.Sleep(10 * time.Millisecond)
	} else {
		time.Sleep(200 * time.Millisecond)
	}
	return 42
}

func withTimeout(timeout time.Duration, fn func() int) (int, error) {
	done := make(chan int, 1)
	go func() {
		res := fn()
		done <- res
		close(done)
	}()
	select {
	case ans := <-done:
		return ans, nil
	case <-time.After(timeout):
		return 0, errors.New("ghee khatam")
	}
}

func main() {
	for range 10 {
		start := time.Now()
		timeout := 50 * time.Millisecond
		if ans, err := withTimeout(timeout, work); err != nil {
			fmt.Printf("took longer than %v. Error: %v\n", time.Since(start), err)
		} else {
			fmt.Printf("took %v. Result: %v\n", time.Since(start), ans)
		}
	}
	// handle, wait := throttle(1, work)
	// start := time.Now()
	// err := handle()
	// fmt.Println("1st call, error:", err)
	//
	// err = handle()
	// fmt.Println("2nd call, error:", err)
	//
	// err = handle()
	// fmt.Println("3rd call, error:", err)
	//
	// err = handle()
	// fmt.Println("4th call, error:", err)
	//
	// wait()
	// fmt.Println("4 calls took", time.Since(start))
}
