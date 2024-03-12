package main

import (
	"slices"
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (43.38%)
 * Likes:    5344
 * Dislikes: 6282
 * Total Accepted:    1.7M
 * Total Submissions: 3.8M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * A phrase is a palindrome if, after converting all uppercase letters into
 * lowercase letters and removing all non-alphanumeric characters, it reads the
 * same forward and backward. Alphanumeric characters include letters and
 * numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "A man, a plan, a canal: Panama"
 * Output: true
 * Explanation: "amanaplanacanalpanama" is a palindrome.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "race a car"
 * Output: false
 * Explanation: "raceacar" is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: s = " "
 * Output: true
 * Explanation: s is an empty string "" after removing non-alphanumeric
 * characters.
 * Since an empty string reads the same forward and backward, it is a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s consists only of printable ASCII characters.
 *
 *
 */
func isPalindrome(s string) bool {
	sl := strings.ToLower(s)
	aslr := slices.DeleteFunc([]rune(sl), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	l, r := 0, len(aslr)-1

	for l < r {
		if aslr[l] != aslr[r] {
			return false
		}
		l++
		r--
	}

	return true

}
