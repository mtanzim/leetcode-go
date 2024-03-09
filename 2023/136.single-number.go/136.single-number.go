package main

import "sort"

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (69.99%)
 * Likes:    11887
 * Dislikes: 452
 * Total Accepted:    1.9M
 * Total Submissions: 2.7M
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers nums, every element appears twice except
 * for one. Find that single one.
 *
 * You must implement a solution with a linear runtime complexity and use only
 * constant extra space.
 *
 *
 * Example 1:
 * Input: nums = [2,2,1]
 * Output: 1
 * Example 2:
 * Input: nums = [4,1,2,1,2]
 * Output: 4
 * Example 3:
 * Input: nums = [1]
 * Output: 1
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -3 * 10^4 <= nums[i] <= 3 * 10^4
 * Each element in the array appears twice except for one element which appears
 * only once.
 *
 *
 */

// @lc code=start
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		if nums[i-1] != nums[i] {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}

// @lc code=end
