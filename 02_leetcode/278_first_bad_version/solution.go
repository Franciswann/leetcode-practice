package first_bad_version_278

/*
 * Problem: 278. First Bad Version
 * Time Spent: 38:04.40
 * Difficulty: Easy
 * Date: 2026-03-24
 */

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	// if n = 5 ->
	// [1, 2, 3, 4, 5]

	right := n
	left := 1

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			// first bad version might be mid or on the left side
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
