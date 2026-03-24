package binary_search_704

/*
 * Problem: 704. Binary Search
 * Time Spent: 27:10.45
 * Difficulty: Easy
 * Date: 2026-03-22
 */

func Binary_search(nums []int, target int) int {
	length := len(nums)
	high := length - 1
	low := 0

	for low <= high {
		// to prevent overflow
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// if target is on the left side, shrink the right boundary
			high = mid - 1
		} else {
			// if target is on the right side, shrink the left boundary
			low = mid + 1
		}
	}
	// target not found, return -1
	return -1
}
