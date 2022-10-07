package main

/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 *
 * https://leetcode.com/problems/first-bad-version/description/
 *
 * algorithms
 * Easy (42.85%)
 * Likes:    6181
 * Dislikes: 2352
 * Total Accepted:    1.2M
 * Total Submissions: 2.9M
 * Testcase Example:  '5\n4'
 *
 * You are a product manager and currently leading a team to develop a new
 * product. Unfortunately, the latest version of your product fails the quality
 * check. Since each version is developed based on the previous version, all
 * the versions after a bad version are also bad.
 *
 * Suppose you have n versions [1, 2, ..., n] and you want to find out the
 * first bad one, which causes all the following ones to be bad.
 *
 * You are given an API bool isBadVersion(version) which returns whether
 * version is bad. Implement a function to find the first bad version. You
 * should minimize the number of calls to the API.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 5, bad = 4
 * Output: 4
 * Explanation:
 * call isBadVersion(3) -> false
 * call isBadVersion(5) -> true
 * call isBadVersion(4) -> true
 * Then 4 is the first bad version.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1, bad = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= bad <= n <= 2^31 - 1
 *
 *
 */

// Simply for testing
func isBadVersion(n int) bool {
	versions := []int{0, 0, 0, 0, 1, 1}
	return versions[n] == 1
}

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 0
	r := n
	for l < r {
		mid := l + (r-l)/2
		isMidBad := isBadVersion(mid)
		if isMidBad {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l

}

// @lc code=end
