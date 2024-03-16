package main

import "slices"

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Medium (48.56%)
 * Likes:    15221
 * Dislikes: 302
 * Total Accepted:    1.4M
 * Total Submissions: 2.9M
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security systems
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given an integer array nums representing the amount of money of each house,
 * return the maximum amount of money you can rob tonight without alerting the
 * police.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 *
 *
 */

func rob(nums []int) int {
	cache := make(map[int]int)

	var traverse func(curIdx int, curNums []int) int
	traverse = func(curIdx int, curNums []int) int {
		if v, ok := cache[curIdx]; ok {
			return v
		}
		if len(curNums) <= 2 {
			return slices.Max(curNums)
		}
		head, neck, rest := curNums[0], curNums[1], curNums[2:]
		willRob := head + traverse(curIdx+2, rest)
		willNotRob := traverse(curIdx+1, append([]int{neck}, rest...))
		curRes := slices.Max([]int{willRob, willNotRob})
		cache[curIdx] = curRes
		return curRes
	}
	return traverse(0, nums)
}
