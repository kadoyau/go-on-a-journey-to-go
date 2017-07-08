package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	splitted := strings.Split(s, " ")
	for i := 0; i < len(splitted); i++ {
		m[splitted[i]] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}