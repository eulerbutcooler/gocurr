package main

import (
	"strings"
	"sync"
	"unicode"
)

func countDigitsInWords(sent string) int {
	var wg sync.WaitGroup
	syncStats := new(sync.Map)
	words := strings.FieldsSeq(sent)

	for r := range words {
		wg.Go(func() {
			count := countDigits(r)
			for {
				prev, loaded := syncStats.LoadOrStore(r, count)
				if !loaded {
					break
				}
				if syncStats.CompareAndSwap(r, prev, prev.(int)+count) {
					break
				}
			}
		})
	}
	wg.Wait()
	stats := asStats(syncStats)
	return stats
}

func asStats(syncStats *sync.Map) int {
	total := 0
	syncStats.Range(func(_, value any) bool {
		total += value.(int)
		return true
	})
	return total
}

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}
