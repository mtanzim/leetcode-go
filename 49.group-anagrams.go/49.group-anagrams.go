package main

import "sort"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (65.81%)
 * Likes:    13275
 * Dislikes: 394
 * Total Accepted:    1.7M
 * Total Submissions: 2.6M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings strs, group the anagrams together. You can return
 * the answer in any order.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * Input: strs = ["a"]
 * Output: [["a"]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lowercase English letters.
 *
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

func groupAnagrams(strs []string) [][]string {

	groups := make(map[string][]string)
	for _, v := range strs {
		runes := []rune(v)
		sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] })
		sortedString := string(runes)
		groups[sortedString] = append(groups[sortedString], v)
	}
	res := make([][]string, len(groups))
	idx := 0
	for _, v := range groups {
		res[idx] = v
		idx++
	}

	return res
}

// @lc code=end
