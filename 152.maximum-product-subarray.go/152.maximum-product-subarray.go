package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (34.86%)
 * Likes:    16196
 * Dislikes: 493
 * Total Accepted:    1M
 * Total Submissions: 2.9M
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find a subarray that has the largest product,
 * and return the product.
 *
 * The test cases are generated so that the answer will fit in a 32-bit
 * integer.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 */

// @lc code=start

// TODO: works but needs DP
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return int(math.Inf(-1))
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		left := nums[0]
		right := nums[1]
		both := left * right
		all := []int{left, right, both}
		sort.Ints(all)
		return all[len(all)-1]
	}

	maxProd := int(math.Inf(-1))
	curProd := 1
	for _, v := range nums {
		curProd *= v
	}
	if curProd > maxProd {
		maxProd = curProd
	}
	for i, v := range nums {
		left := maxProduct(nums[:i])
		if left > int(maxProd) {
			maxProd = left
		}
		right := maxProduct(nums[i+1:])
		if right > int(maxProd) {
			maxProd = right
		}
		me := maxProduct([]int{v})
		if me > int(maxProd) {
			maxProd = me
		}

	}
	return maxProd
}

// @lc code=end
