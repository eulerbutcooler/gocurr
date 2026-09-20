package main

import (
	"strings"
	"sync"
)

func countDigitsInWordsCh(sent string) int {
	digits := make(chan int)
	go func() {
		var wg sync.WaitGroup
		words := strings.FieldsSeq(sent)
		for r := range words {
			wg.Go(func() {
				digits <- countDigits(r)
			})
		}
		wg.Wait()
		close(digits)
	}()
	total := 0
	for n := range digits {
		total += n
	}
	return total
}
