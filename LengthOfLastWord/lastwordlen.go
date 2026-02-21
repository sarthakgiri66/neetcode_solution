package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1

	// Skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// Count the last word's characters
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))           // Output: 5
	fmt.Println(lengthOfLastWord("   fly me   "))          // Output: 2
	fmt.Println(lengthOfLastWord("luffy is still joyboy")) // Output: 6
}



#This solution is slightly more efficient in practice — it scans from the lef and stops as soon as it counts the space. This solution always reads partial(stops early) in the traversal
