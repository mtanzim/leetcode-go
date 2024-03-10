package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (40.72%)
 * Likes:    19541
 * Dislikes: 1137
 * Total Accepted:    3.3M
 * Total Submissions: 8.1M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: s = "(]"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, c := range s {
		// it's an opener
		if _, ok := pairs[c]; ok {
			stack = append(stack, c)
			continue
		} else {
			if len(stack) == 0 {
				return false
			}
		}

		if len(stack) == 0 {
			continue
		}

		// it's a closer
		top := stack[len(stack)-1]
		if c == pairs[top] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}

	return len(stack) == 0
}

// @lc code=end
