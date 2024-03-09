package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (45.82%)
 * Likes:    19128
 * Dislikes: 647
 * Total Accepted:    1.9M
 * Total Submissions: 4.1M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given an array of intervals where intervals[i] = [starti, endi], merge all
 * overlapping intervals, and return an array of the non-overlapping intervals
 * that cover all the intervals in the input.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlap, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// sort by start
	// sortedIntervals := copy([][]int{}, intervals)

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 1
	// consume greedily
	for i < len(intervals) {
		firstEnd, secondStart := intervals[i-1][1], intervals[i][0]
		firstStart, secondEnd := intervals[i-1][0], intervals[i][1]
		if secondStart <= firstEnd {
			minStart, maxEnd := firstStart, secondEnd
			if secondStart < firstStart {
				minStart = secondStart
			}
			if firstEnd > secondEnd {
				maxEnd = firstEnd
			}
			intervals[i] = []int{minStart, maxEnd}
			intervals[i-1] = []int{}
		}
		i += 1
	}
	// filter out
	res := [][]int{}
	for _, v := range intervals {
		if len(v) != 0 {
			res = append(res, v)
		}
	}

	return res

}

// @lc code=end
