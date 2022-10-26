package main

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (61.22%)
 * Likes:    7066
 * Dislikes: 1041
 * Total Accepted:    2.2M
 * Total Submissions: 3.6M
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an integer array nums, return true if any value appears at least twice
 * in the array, and return false if every element is distinct.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,1]
 * Output: true
 * Example 2:
 * Input: nums = [1,2,3,4]
 * Output: false
 * Example 3:
 * Input: nums = [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	hm := make(map[int]struct{})
	for _, v := range nums {
		_, ok := hm[v]
		if ok {
			return true
		}
		hm[v] = struct{}{}
	}
	return false
}

// @lc code=end
