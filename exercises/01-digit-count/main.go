package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sent := "meo1 mai3 hue23 catika1234189324uasba321 w32uq78321"
	if len(os.Args) > 1 {
		sent = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("input: %q\n", sent)
	for _, im := range impls {
		fmt.Printf("  %-16s %d\n", im.name, im.fn(sent))
	}
}
