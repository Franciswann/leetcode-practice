package main

import "fmt"

func Binary_search(arr []int, item int) int {
	high := len(arr) - 1
	low := 0
	mid := (high + low) / 2

	for i := 1; i < len(arr); i++ {

		guess := arr[mid]

		// guess correct
		if guess == item {
			return mid

			// guess too high
		} else if guess > item {
			high = mid - 1
			mid = (high + low) / 2

			// guess too low
		} else {
			low = mid + 1
			mid = (high + low) / 2
		}
	}
	// return -1 if the item is not found
	return -1
}

func main() {
	// index := [0, 1, 2, 3, 4]
	arr := []int{1, 3, 5, 7, 9}
	result := Binary_search(arr, 9)
	fmt.Println("result: ", result)
}
