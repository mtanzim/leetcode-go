package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (33.74%)
 * Likes:    33648
 * Dislikes: 1468
 * Total Accepted:    4.5M
 * Total Submissions: 13.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string s, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * Notice that the answer must be a substring, "pwke" is a subsequence and not
 * a substring.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s consists of English letters, digits, symbols and spaces.
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	curHM := make(map[rune]struct{})
	l := 0
	r := 0
	max := 0
	for r < len(s) {
		cur := s[r]
		if _, ok := curHM[rune(cur)]; ok {
			curHM = make(map[rune]struct{})
			delta := r - l
			l++
			r = l
			if delta > max {
				max = delta
			}
		} else {
			curHM[rune(cur)] = struct{}{}
			r++
		}
	}
	delta := r - l
	if delta > max {
		max = delta
	}
	return max

}

// @lc code=end
