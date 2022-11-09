package main

import (
	"fmt"
	"strings"
)

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

// @lc code=start

func max(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func intSliceToStr(nums []int) string {
	var sb strings.Builder
	for _, v := range nums {
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	return sb.String()
}

func robCached(nums []int, cache map[string]int) int {
	if len(nums) <= 2 {
		return max(nums)
	}
	maxAmt := 0
	for i := 0; i < len(nums)-1; i++ {
		s := nums[i]
		remaining := nums[i+2:]
		remainingKey := intSliceToStr(nums)
		var canRobWithS int
		if robCachedVal, ok := cache[remainingKey]; ok {
			canRobWithS = s + robCachedVal
		} else {
			canRobWithS = s + robCached(remaining, cache)
			cache[remainingKey] = canRobWithS
		}
		maxAmt = max([]int{maxAmt, canRobWithS})
	}
	return maxAmt
}

func rob(nums []int) int {
	cache := make(map[string]int)
	return robCached(nums, cache)
}

// @lc code=end
