package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (74.50%)
 * Likes:    13552
 * Dislikes: 226
 * Total Accepted:    1.4M
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,2,3]'
 *
 * Given an array nums of distinct integers, return all the possible
 * permutations. You can return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * Example 2:
 * Input: nums = [0,1]
 * Output: [[0,1],[1,0]]
 * Example 3:
 * Input: nums = [1]
 * Output: [[1]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * All the integers of nums are unique.
 *
 *
 */

// @lc code=start
func permute(nums []int) [][]int {

	if len(nums) == 1 {
		a := nums[0]
		return [][]int{[]int{a}}
	}

	if len(nums) > 1 {
		a := nums[0]
		rest := permute(nums[1:])
		results := [][]int{}
		for _, curArr := range rest {
			for i := 0; i < len(curArr); i++ {
				// left and in between
				newEntry := append([]int{}, curArr[:i]...)
				newEntry = append(newEntry, a)
				newEntry = append(newEntry, curArr[i:]...)
				results = append(results, newEntry)
			}
			// right
			newEntryRight := append([]int{}, curArr...)
			newEntryRight = append(newEntryRight, a)
			results = append(results, newEntryRight)

		}

		return results
	}
	return [][]int{}

}

// @lc code=end
