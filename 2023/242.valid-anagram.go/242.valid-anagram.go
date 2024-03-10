package main

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (62.66%)
 * Likes:    7508
 * Dislikes: 248
 * Total Accepted:    1.8M
 * Total Submissions: 2.9M
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t, return true if t is an anagram of s, and false
 * otherwise.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * Example 2:
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, t.length <= 5 * 10^4
 * s and t consist of lowercase English letters.
 *
 *
 *
 * Follow up: What if the inputs contain Unicode characters? How would you
 * adapt your solution to such a case?
 *
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	hmS := make(map[rune]int)
	hmT := make(map[rune]int)
	for _, v := range s {
		hmS[v]++
	}
	for _, v := range t {
		hmT[v]++
	}
	for _, v := range s {
		if hmT[v] != hmS[v] {
			return false
		}
	}
	for _, v := range t {
		if hmT[v] != hmS[v] {
			return false
		}
	}
	return true
}

// @lc code=end
