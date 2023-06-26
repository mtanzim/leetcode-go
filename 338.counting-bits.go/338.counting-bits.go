package main

import "strconv"

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 *
 * https://leetcode.com/problems/counting-bits/description/
 *
 * algorithms
 * Easy (75.92%)
 * Likes:    8961
 * Dislikes: 432
 * Total Accepted:    789K
 * Total Submissions: 1M
 * Testcase Example:  '2'
 *
 * Given an integer n, return an array ans of length n + 1 such that for each i
 * (0 <= i <= n), ans[i] is the number of 1's in the binary representation of
 * i.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: [0,1,1]
 * Explanation:
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 *
 *
 * Example 2:
 *
 *
 * Input: n = 5
 * Output: [0,1,1,2,1,2]
 * Explanation:
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 * 3 --> 11
 * 4 --> 100
 * 5 --> 101
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 10^5
 *
 *
 *
 * Follow up:
 *
 *
 * It is very easy to come up with a solution with a runtime of O(n log n). Can
 * you do it in linear time O(n) and possibly in a single pass?
 * Can you do it without using any built-in function (i.e., like
 * __builtin_popcount in C++)?
 *
 *
 */

// @lc code=start
func hammingWeightStringManipulation(num uint32) int {
	binary := strconv.FormatUint(uint64(num), 2)
	count := 0
	for _, c := range binary {
		if c == '1' {
			count++
		}
	}
	return count
}

func hammingWeight(num uint32) int {

	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count

}

func countBits(n int) []int {
	ans := []int{}
	for i := 0; i <= n; i++ {
		ans = append(ans, hammingWeight(uint32(i)))
	}
	return ans
}

// @lc code=end
