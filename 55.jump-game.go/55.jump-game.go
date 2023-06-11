package main

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 *
 * https://leetcode.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (38.40%)
 * Likes:    16355
 * Dislikes: 850
 * Total Accepted:    1.4M
 * Total Submissions: 3.6M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * You are given an integer array nums. You are initially positioned at the
 * array's first index, and each element in the array represents your maximum
 * jump length at that position.
 *
 * Return true if you can reach the last index, or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last
 * index.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its maximum
 * jump length is 0, which makes it impossible to reach the last index.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 10^5
 *
 *
 */

// @lc code=start

func traverse(nums []int, pos, target int) bool {
	if pos == target {
		return true
	}
	if pos > target {
		return false
	}
	curJump := nums[pos]
	canJump := false
	for i := 1; i <= curJump; i++ {
		canJump = canJump || traverse(nums, pos+i, target)
	}
	return canJump
}

func canJump(nums []int) bool {
	target := len(nums) - 1
	return traverse(nums, 0, target)
}

// @lc code=end
