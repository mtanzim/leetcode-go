package main

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (58.77%)
 * Likes:    6886
 * Dislikes: 233
 * Total Accepted:    1.3M
 * Total Submissions: 2.2M
 * Testcase Example:  '"leetcode"'
 *
 * Given a string s, find the first non-repeating character in it and return
 * its index. If it does not exist, return -1.
 *
 *
 * Example 1:
 * Input: s = "leetcode"
 * Output: 0
 * Example 2:
 * Input: s = "loveleetcode"
 * Output: 2
 * Example 3:
 * Input: s = "aabb"
 * Output: -1
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^5
 * s consists of only lowercase English letters.
 *
 *
 */

// @lc code=start
func firstUniqChar(s string) int {
	hm := make(map[rune]int)
	for _, c := range s {
		hm[c]++
	}
	for i, v := range s {
		if hm[v] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
