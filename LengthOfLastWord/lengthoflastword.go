package main

import "fmt"

func lengthOfLastWord(inp string) int {
  lastWordLen := 0
  letterCount := 0 
  var prevChar rune

  for _, currentChar := range inp {
    if unicode.IsSpace(currentChar) {
      if unicode.IsSpace(prevChar) == false {
        lastWordLen = letterCount
      }
      letterCount = 0 
    } else {
      letterCount++
    }
    prevChar = currentChar
  }

  if letterCount > 0 {
    lastWordLen = letterCount
  }

  return lastWordLen
}


func main() {
	fmt.Println(lengthOfLastWord("Hello World"))           // Output: 5
	fmt.Println(lengthOfLastWord("   fly me   "))          // Output: 2
	fmt.Println(lengthOfLastWord("luffy is still joyboy")) // Output: 6
}
