package two_sum_001

/*
 * Problem: 704. Binary Search
 * Time Spent: 40:08.51
 * Difficulty: Easy
 * Date: 2026-04-02
 */

func twoSum(nums []int, target int) []int {

	List := make(map[int]int)

	for i, num := range nums {

		difference := target - num

		if v, ok := List[difference]; ok {
			return []int{v, i}
		}

		List[num] = i
	}
	return nil
}
