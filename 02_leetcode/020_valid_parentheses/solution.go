package valid_parentheses_020

/*
 * Problem: 20. Valid Parentheses
 * Time Spent: 1:08:14
 * Difficulty: Easy
 * Date: 2026-04-05
 */

func isValid(s string) bool {
	List := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, i := range s {
		if i == '(' || i == '[' || i == '{' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != List[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
