package main

import (
	"strings"
	"sync"
	"sync/atomic"
)

func countDigitsInWordsAt(sent string) int {
	var wg sync.WaitGroup
	var totaldigits atomic.Uint64
	words := strings.FieldsSeq(sent)

	for r := range words {
		wg.Go(func() {
			totaldigits.Add(uint64(countDigits(r)))
		})
	}
	wg.Wait()
	return int(totaldigits.Load())
}
