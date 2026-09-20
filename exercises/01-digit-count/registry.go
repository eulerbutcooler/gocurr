package main

type impl struct {
	name string
	fn   func(string) int
}

var impls = []impl{
	{"v1_syncmap", countDigitsInWords},
	{"v2_channel", countDigitsInWordsCh},
	{"v3_atomic", countDigitsInWordsAt},
}
