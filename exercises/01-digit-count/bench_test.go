package main

import (
	"strings"
	"testing"
)

func BenchmarkCountDigitsInWords(b *testing.B) {
	benches := []struct {
		size string
		sent string
	}{
		{"small", "meo1 mai2 hu catik21123a"},
		{"large", strings.Repeat("meo1 mai2 hu catik21123a ", 2000)},
	}
	for _, bb := range benches {
		for _, im := range impls {
			b.Run(bb.size+"/"+im.name, func(b *testing.B) {
				sent := bb.sent
				b.ReportAllocs()
				for b.Loop() {
					im.fn(sent)
				}
			})
		}
	}
}

func BenchmarkCountDigits(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		countDigits("catik21123a")
	}
}
