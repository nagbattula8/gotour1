package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	fields := strings.Fields(s)

	dict := make(map[string]int)

	for i := 0; i < len(fields); i++ {

		dict[fields[i]] += 1
	}

	return dict

}

func main() {
	wc.Test(WordCount)
}
