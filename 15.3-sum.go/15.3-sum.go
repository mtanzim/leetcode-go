package main

import (
	"fmt"
	"sort"
	"strings"
)

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
func twoSum(nums []int, target int) [][]int {
	hm := make(map[int]int)
	for i, v := range nums {
		hm[v] = i
	}

	combos := [][]int{}

	for i := range nums {
		remaining := target - nums[i]
		if j, ok := hm[remaining]; ok && j != i {
			combos = append(combos, []int{i, j})
		}
	}
	return combos
}

func arrToStr(a []int) string {
	var sb strings.Builder
	for _, v := range a {
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	return sb.String()
}

func threeSum(nums []int) [][]int {

	combos := [][]int{}
	combosHM := make(map[string]struct{})

	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		twoSumVals := twoSum(nums, target)
		for _, twoSumVal := range twoSumVals {
			if len(twoSumVal) == 2 {
				j := twoSumVal[0]
				k := twoSumVal[1]

				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					newEntry := []int{nums[i], nums[j], nums[k]}
					sort.Ints(newEntry)
					if _, ok := combosHM[arrToStr(newEntry)]; !ok {
						combos = append(combos, newEntry)
						combosHM[arrToStr(newEntry)] = struct{}{}
					}
				}
			}
		}
	}
	return combos
}

// @lc code=end
