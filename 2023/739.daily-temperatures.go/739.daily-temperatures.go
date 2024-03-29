package main

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 *
 * https://leetcode.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (66.81%)
 * Likes:    10520
 * Dislikes: 237
 * Total Accepted:    603.7K
 * Total Submissions: 911.2K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * Given an array of integers temperatures represents the daily temperatures,
 * return an array answer such that answer[i] is the number of days you have to
 * wait after the i^th day to get a warmer temperature. If there is no future
 * day for which this is possible, keep answer[i] == 0 instead.
 *
 *
 * Example 1:
 * Input: temperatures = [73,74,75,71,69,72,76,73]
 * Output: [1,1,4,2,1,1,0,0]
 * Example 2:
 * Input: temperatures = [30,40,50,60]
 * Output: [1,1,1,0]
 * Example 3:
 * Input: temperatures = [30,60,90]
 * Output: [1,1,0]
 *
 *
 * Constraints:
 *
 *
 * 1 <= temperatures.length <= 10^5
 * 30 <= temperatures[i] <= 100
 *
 *
 */

// @lc code=start

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		j := i + 1

		for j < len(temperatures) && temperatures[j] <= temperatures[i] {
			if result[j] <= 0 {
				break
			}
			j += result[j]
		}

		if j < len(temperatures) && temperatures[j] > temperatures[i] {
			result[i] = j - i
		}

	}
	return result
}

// @lc code=end
