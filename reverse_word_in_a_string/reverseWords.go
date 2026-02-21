package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	words := strings.Fields(s)
	i := 0
	j := len(words) - 1

	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	out := strings.Join(words, " ")
	return out
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))  // "blue is sky the"
	fmt.Println(reverseWords("  hello world  "))  // "world hello"
	fmt.Println(reverseWords("a good   example")) // "example good a"
}

#It's O(n) time, O(n) space
