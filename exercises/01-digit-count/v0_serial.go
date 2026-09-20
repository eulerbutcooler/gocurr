package main

import "strings"

func countDigitsInWordsSerial(str string) int {
	total := 0
	for word := range strings.FieldsSeq(str) {
		total += countDigits(word)
	}
	return total
}
