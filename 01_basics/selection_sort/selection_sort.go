package main

import (
	"fmt"
)

// // 1. create empty array and put smallest number into it
// func FindSmallest(arr []int) int {

// 	number := arr[0]

// 	for index, _ := range arr {
// 		if arr[index] < number {
// 			number = arr[index]
// 		}
// 	}
// 	// index := slices.Index(arr, number)
// 	// slices.Delete(arr, index, index+1)
// 	// fmt.Println("array: ", arr)
// 	return number
// }

// func SelectSort(arr []int) []int {
// 	result := make([]int, len(arr))

// 	for index, _ := range arr {
// 		result[index] = FindSmallest(arr)
// 	}
// 	return result
// }

// 2. exchange smallest number in current array
func SelectSort_2(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIdx := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	result := SelectSort_2(arr)
	fmt.Println(result)
}
