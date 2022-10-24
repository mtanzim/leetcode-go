package main

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	// try O(n^2 solution first)
	k := len(nums)
	for i := 0; i < k-1; i++ {
		j := i + 1
		for j < k && j < len(nums) && nums[i] == nums[j] {
			j++
		}
		delta := j - i - 1
		if delta > 0 {
			for m := i + 1; m+delta < k && m+delta < len(nums); m++ {
				nums[m] = nums[m+delta]
			}
			k -= delta
		}
	}
	return k
}

// @lc code=end
