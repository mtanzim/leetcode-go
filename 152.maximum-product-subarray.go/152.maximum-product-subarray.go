package main

import "math"

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

func subarrays(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	if len(nums) == 2 {
		left := []int{nums[0]}
		right := []int{nums[1]}
		return [][]int{left, nums, right}
	}

	overallRes := [][]int{}
	for i, v := range nums {
		left := subarrays(nums[:i])
		right := subarrays(nums[i+1:])
		me := subarrays([]int{v})
		overallRes = append(overallRes, left...)
		overallRes = append(overallRes, right...)
		overallRes = append(overallRes, me...)
	}
	return overallRes

}

func maxProduct(nums []int) int {
	arrs := subarrays(nums)
	maxProd := math.Inf(-1)
	for _, arr := range arrs {
		prod := 1.0
		for _, v := range arr {
			prod *= float64(v) 
		}
		if prod > maxProd {
			maxProd = float64(prod)
		}
	}
	return int(maxProd)
}

// @lc code=end
