package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Medium (50.01%)
 * Likes:    26244
 * Dislikes: 1180
 * Total Accepted:    2.8M
 * Total Submissions: 5.7M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the subarray which has the largest sum and
 * return its sum.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1]
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [5,4,-1,7,8]
 * Output: 23
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 * Follow up: If you have figured out the O(n) solution, try coding another
 * solution using the divide and conquer approach, which is more subtle.
 *
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		a, b := nums[0], nums[1]
		if a+b > a && a+b > b {
			return a + b
		}
		if a > b {
			return a
		}
		return b

	}

	maxSum := 0
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		rest := append([]int{}, nums[0:i]...)
		rest = append(rest, nums[i+1:]...)
		restSum := maxSubArray(rest)
		if cur+restSum > cur && cur+restSum > restSum && cur+restSum > maxSum {
			maxSum = cur + restSum
		} else if cur > restSum && cur > maxSum {
			maxSum = cur
		} else if restSum > maxSum {
			maxSum = restSum
		}
	}

	return maxSum

}

// @lc code=end
