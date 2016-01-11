package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) (result map[string]int) {
	result = make(map[string]int)
	var tokenizedString []string = strings.Fields(s)
	for _, value := range tokenizedString {
		if count, ok := result[value]; ok {
			result[value] = count+1
		} else {
			result[value] = 1
		}
	}
	return
}

func main() {
	wc.Test(WordCount)
}

