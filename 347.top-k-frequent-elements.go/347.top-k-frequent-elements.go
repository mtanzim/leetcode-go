package main

import "sort"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (64.84%)
 * Likes:    13001
 * Dislikes: 477
 * Total Accepted:    1.4M
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given an integer array nums and an integer k, return the k most frequent
 * elements. You may return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 * Example 2:
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * k is in the range [1, the number of unique elements in the array].
 * It is guaranteed that the answer is unique.
 *
 *
 *
 * Follow up: Your algorithm's time complexity must be better than O(n log n),
 * where n is the array's size.
 *
 */

// @lc code=start
type container struct {
	v int
	n int
}

func topKFrequent(nums []int, k int) []int {

	hm := make(map[int]int)
	for _, v := range nums {
		hm[v]++
	}

	cs := []container{}
	for k, v := range hm {
		cs = append(cs, container{k, v})
	}
	sort.Slice(cs, func(i, j int) bool { return cs[i].n > cs[j].n })

	r := make([]int,k)

	for i := range r {
		r[i] = cs[i].v
	}
	return r

}

// @lc code=end
