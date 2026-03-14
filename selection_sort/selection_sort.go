package main

import (
	"fmt"
)

func FindSmallest(arr []int) int {

	number := arr[0]

	for index, _ := range arr {
		if arr[index] < number {
			number = arr[index]
		}
	}
	// index := slices.Index(arr, number)
	// slices.Delete(arr, index, index+1)
	// fmt.Println("array: ", arr)
	return number
}

func SelectSort(arr []int) []int {
	result := make([]int, len(arr))

	for index, _ := range arr {
		result[index] = FindSmallest(arr)
	}
	return result
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	result := SelectSort(arr)
	fmt.Println(result)
}
