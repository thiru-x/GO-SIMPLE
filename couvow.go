package main

import (
	"fmt"
	"strings"
)

func countVowelsConsonants(s string) (int, int) {
	s = strings.ToLower(s)
	vowels, consonants := 0, 0

	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l',
			'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
			consonants++
		}
	}
	return vowels, consonants
}
func main() {
	// Example usage
	input := "Hello, World!"
	vowels, consonants := countVowelsConsonants(input)
	fmt.Println("Counting vowels and consonants in Go")
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Vowels: %d, Consonants: %d\n", vowels, consonants)
}
