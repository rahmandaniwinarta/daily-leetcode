package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
}

func removeOccurrences(s string, part string) string {
	iteration := 1
	for strings.Contains(s, part) {
		fmt.Printf("Iteration %d: Before removing '%s' -> %s\n", iteration, part, s)
		s = strings.Replace(s, part, "", 1)
		fmt.Printf("Iteration %d: After removing  '%s' -> %s\n\n", iteration, part, s)
		iteration++
	}
	return string(s)
}
