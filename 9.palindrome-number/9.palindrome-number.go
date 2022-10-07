package main

import "strconv"

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (52.79%)
 * Likes:    7443
 * Dislikes: 2318
 * Total Accepted:    2.5M
 * Total Submissions: 4.8M
 * Testcase Example:  '121'
 *
 * Given an integer x, return true if x is palindrome integer.
 *
 * An integer is a palindrome when it reads the same backward as forward.
 *
 *
 * For example, 121 is a palindrome while 123 is not.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: x = 121
 * Output: true
 * Explanation: 121 reads as 121 from left to right and from right to left.
 *
 *
 * Example 2:
 *
 *
 * Input: x = -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: x = 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you solve it without converting the integer to a string?
 */

// @lc code=start
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	runes := []rune(s)
	i := 0
	j := len(runes) - 1
	for i < len(runes) && j > 0 && j > i {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end
