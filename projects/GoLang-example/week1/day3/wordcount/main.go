package main

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	curr := strings.Fields(str)
	counts := map[string]int{}
	for _, mywords := range curr {
		counts[mywords]++
	}
	return counts
}

func main() {
	str := "Anything I want this message to Be like Test Test"
	result := WordCount(str)
	fmt.Println(result)
}
