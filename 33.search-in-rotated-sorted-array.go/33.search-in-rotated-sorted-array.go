package main

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (38.57%)
 * Likes:    21093
 * Dislikes: 1260
 * Total Accepted:    2M
 * Total Submissions: 5.2M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * There is an integer array nums sorted in ascending order (with distinct
 * values).
 *
 * Prior to being passed to your function, nums is possibly rotated at an
 * unknown pivot index k (1 <= k < nums.length) such that the resulting array
 * is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
 * (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3
 * and become [4,5,6,7,0,1,2].
 *
 * Given the array nums after the possible rotation and an integer target,
 * return the index of target if it is in nums, or -1 if it is not in nums.
 *
 * You must write an algorithm with O(log n) runtime complexity.
 *
 *
 * Example 1:
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * Example 2:
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * Example 3:
 * Input: nums = [1], target = 0
 * Output: -1
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * All values of nums are unique.
 * nums is an ascending array that is possibly rotated.
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start

func findMinIdx(nums []int) int {
	var traverse func(low, high int, arr []int) (bool, int)
	traverse = func(low, high int, arr []int) (bool, int) {
		if low > high {
			return false, -1
		}
		mid := (low + high) / 2
		if mid > 0 && arr[mid] < arr[mid-1] {
			return true, mid
		}
		if mid < len(arr)-1 && arr[mid] > arr[mid+1] {
			return true, mid + 1
		}
		leftOk, leftIdx := traverse(mid+1, high, arr)
		rightOk, rightIdx := traverse(low, mid-1, arr)
		if leftOk {
			return leftOk, leftIdx
		}
		if rightOk {
			return rightOk, rightIdx
		}
		return false, -1
	}

	if len(nums) == 1 {
		return nums[0]
	}
	ok, idx := traverse(0, len(nums)-1, nums)
	if !ok {
		return 0
	}
	return idx

}

func binarySearch(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	curVal := nums[mid]
	if curVal == target {
		return mid
	}
	if target < curVal {
		return binarySearch(nums, target, lo, mid-1)
	}
	return binarySearch(nums, target, mid+1, hi)

}

func search(nums []int, target int) int {
	startingIdx := findMinIdx(nums)
	newArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newArr[i] = nums[(i+startingIdx)%len(nums)]
	}
	newIdx := binarySearch(newArr, target, 0, len(nums)-1)
	if newIdx != -1 {
		return newIdx + startingIdx
	}
	return -1
}

// @lc code=end
