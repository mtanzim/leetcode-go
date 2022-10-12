package main

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 *
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (59.54%)
 * Likes:    7459
 * Dislikes: 411
 * Total Accepted:    647.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array nums of n integers where nums[i] is in the range [1, n],
 * return an array of all the integers in the range [1, n] that do not appear
 * in nums.
 *
 *
 * Example 1:
 * Input: nums = [4,3,2,7,8,2,3,1]
 * Output: [5,6]
 * Example 2:
 * Input: nums = [1,1]
 * Output: [2]
 *
 *
 * Constraints:
 *
 *
 * n == nums.length
 * 1 <= n <= 10^5
 * 1 <= nums[i] <= n
 *
 *
 *
 * Follow up: Could you do it without extra space and in O(n) runtime? You may
 * assume the returned list does not count as extra space.
 *
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	tracker := []int{}
	hm := make(map[int]bool)
	for _, v := range nums {
		hm[v] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !hm[i] {
			tracker = append(tracker, i)
		}
	}
	return tracker

}

// @lc code=end
