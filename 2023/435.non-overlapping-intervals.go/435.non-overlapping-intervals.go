package main

import "sort"

/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 *
 * https://leetcode.com/problems/non-overlapping-intervals/description/
 *
 * algorithms
 * Medium (49.69%)
 * Likes:    5889
 * Dislikes: 161
 * Total Accepted:    368.6K
 * Total Submissions: 730.9K
 * Testcase Example:  '[[1,2],[2,3],[3,4],[1,3]]'
 *
 * Given an array of intervals intervals where intervals[i] = [starti, endi],
 * return the minimum number of intervals you need to remove to make the rest
 * of the intervals non-overlapping.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
 * Output: 1
 * Explanation: [1,3] can be removed and the rest of the intervals are
 * non-overlapping.
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,2],[1,2],[1,2]]
 * Output: 2
 * Explanation: You need to remove two [1,2] to make the rest of the intervals
 * non-overlapping.
 *
 *
 * Example 3:
 *
 *
 * Input: intervals = [[1,2],[2,3]]
 * Output: 0
 * Explanation: You don't need to remove any of the intervals since they're
 * already non-overlapping.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * -5 * 10^4 <= starti < endi <= 5 * 10^4
 *
 *
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {

	// sort the intervals by starting
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 1
	count := 0
	for i < len(intervals) {
		aStart, aEnd := intervals[i-1][0], intervals[i-1][1]
		bStart, bEnd := intervals[i][0], intervals[i][1]
		// if they overlap
		if aEnd > bStart {
			count++
			// remove the interval with the widest end
			selectedInterval := []int{aStart, aEnd}
			if bEnd < aEnd {
				selectedInterval = []int{bStart, bEnd}
			}
			intervals[i] = selectedInterval
			intervals[i-1] = []int{}
		}
		i++
	}
	return count

}

// @lc code=end
