package main

import (
    "fmt"
)

func subsets(nums []int) [][]int {
    result := [][]int{{}} // start with empty set

    for _, num := range nums {
        size := len(result)
        for i := 0; i < size; i++ {
            set := append([]int{}, result[i]...)
            set = append(set, num)
            result = append(result, set)
        }
    }
    return result
}

func main() {
    nums := []int{1, 2}
    result := subsets(nums)
    fmt.Println(result)
}
