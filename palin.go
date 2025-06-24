package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	// Example usage
	input := "radar"
	if isPalindrome(input) {
		fmt.Println("Checking if a string is a palindrome in Go")
		fmt.Printf("The string '%s' is a palindrome.\n", input)
	} else {
		fmt.Printf("The string '%s' is not a palindrome.\n", input)
	}
}
