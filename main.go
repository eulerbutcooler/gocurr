package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

func main() {
	countDigitsInWords("meo1 mai2 hu catik21123a")
}

func countDigitsInWords(sent string) int {
	var wg sync.WaitGroup
	syncStats := new(sync.Map)
	words := strings.FieldsSeq(sent)

	for r := range words {
		wg.Go(func() {
			count := countDigits(r)
			syncStats.Store(r, count)
		})
	}
	wg.Wait()
	fmt.Println(asStats(syncStats))
	return asStats(syncStats)
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
