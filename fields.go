package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var res = make(map[string]int)
	for _, val := range strings.Fields(s) {
		res[val] += 1
	}
	return res
}

func field() {
	wc.Test(WordCount)
}
