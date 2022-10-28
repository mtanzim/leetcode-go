package main

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (32.16%)
 * Likes:    22216
 * Dislikes: 2033
 * Total Accepted:    2.3M
 * Total Submissions: 7.2M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an integer array nums, return all the triplets [nums[i], nums[j],
 * nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] +
 * nums[k] == 0.
 *
 * Notice that the solution set must not contain duplicate triplets.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Explanation:
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
 * The distinct triplets are [-1,0,1] and [-1,-1,2].
 * Notice that the order of the output and the order of the triplets does not
 * matter.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,1]
 * Output: []
 * Explanation: The only possible triplet does not sum up to 0.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [0,0,0]
 * Output: [[0,0,0]]
 * Explanation: The only possible triplet sums up to 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, v := range nums {
		hm[v] = i
	}

	for i := range nums {
		remaining := target - nums[i]
		if j, ok := hm[remaining]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}

func threeSum(nums []int) [][]int {

	combos := [][]int{}
	for i := 0; i < len(nums); i++ {
		outer := nums[i]
		twoSumVal := twoSum(nums, outer)
		if len(twoSumVal) == 2 {
			j := twoSumVal[0]
			k := twoSumVal[1]

			if outer+nums[j]+nums[k] == 0 && i != j && i != k {
				combos = append(combos, []int{outer, nums[j], nums[k]})
			}
		}

	}
	return combos
}

// @lc code=end
