package main
import (
	"fmt")
func bubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

func main() {
	nums := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", nums)
	sortedNums := bubbleSort(nums)
	fmt.Println("Sorted array:", sortedNums)
}
