package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (64.63%)
 * Likes:    15438
 * Dislikes: 864
 * Total Accepted:    1.4M
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an integer array nums, return an array answer such that answer[i] is
 * equal to the product of all the elements of nums except nums[i].
 *
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 * You must write an algorithm that runs in O(n) time and without using the
 * division operation.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 *
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The
 * output array does not count as extra space for space complexity analysis.)
 *
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	productHM := make(map[int][]int)
	for i := range nums {
		left := nums[:i]
		right := nums[i+1:]
		rest := append([]int{},left...)
		rest = append(rest,right...)
		productHM[i] = rest
	}
	sumArray := make([]int, len(nums))
	for k, v := range productHM {
		curProd := 1
		for _, val := range v {
			curProd *= val
		}
		sumArray[k] = curProd
	}
	return sumArray
}

// @lc code=end
