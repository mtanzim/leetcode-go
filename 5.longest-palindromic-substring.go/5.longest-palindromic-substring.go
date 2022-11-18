package main

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (32.39%)
 * Likes:    22632
 * Dislikes: 1308
 * Total Accepted:    2.2M
 * Total Submissions: 6.8M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, return the longest palindromic substring in s.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "babad"
 * Output: "bab"
 * Explanation: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "cbbd"
 * Output: "bb"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consist of only digits and English letters.
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	front := 0
	back := len(s) - 1
	for front < back {
		if s[front] != s[back] {
			return false
		}
		front++
		back--
	}
	return true
}

func longestPalindrome(s string) string {
	seen := make(map[rune]int)
	palindromes := []string{}
	for i, v := range []rune(s) {
		if prevLetterIdx, ok := seen[v]; ok {
			runesToCheck := []rune(s)[prevLetterIdx:i+1]
			if isPalindrome(string(runesToCheck)) {
				palindromes = append(palindromes, string(runesToCheck))
			}
		}
		seen[v] = i
	}

	if len(palindromes) == 0 {
		return ""
	}
	longestPalindrome := palindromes[0]
	for _, v := range palindromes {
		if len(v) > len(longestPalindrome) {
			longestPalindrome = v
		}
	}

	return longestPalindrome
}

// @lc code=end
