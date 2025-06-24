package main

import (
	"fmt"
)

func findDuplicates(arr []int) []int {
	count := make(map[int]int)
	var result []int

	for _, num := range arr {
		count[num]++
		if count[num] == 2 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1, 2, 6, 6, 7, 8, 9, 10, 10}
	duplicates := findDuplicates(nums)
	if len(duplicates) > 0 {
		fmt.Println("Duplicates found:", duplicates)
	} else {
		fmt.Println("No duplicates found.")
	}
}
