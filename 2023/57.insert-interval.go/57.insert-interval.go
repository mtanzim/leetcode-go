package main

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 *
 * https://leetcode.com/problems/insert-interval/description/
 *
 * algorithms
 * Medium (37.89%)
 * Likes:    8393
 * Dislikes: 587
 * Total Accepted:    772.1K
 * Total Submissions: 2M
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * You are given an array of non-overlapping intervals intervals where
 * intervals[i] = [starti, endi] represent the start and the end of the i^th
 * interval and intervals is sorted in ascending order by starti. You are also
 * given an interval newInterval = [start, end] that represents the start and
 * end of another interval.
 *
 * Insert newInterval into intervals such that intervals is still sorted in
 * ascending order by starti and intervals still does not have any overlapping
 * intervals (merge overlapping intervals if necessary).
 *
 * Return intervals after the insertion.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * Output: [[1,5],[6,9]]
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * Output: [[1,2],[3,10],[12,16]]
 * Explanation: Because the new interval [4,8] overlaps with
 * [3,5],[6,7],[8,10].
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^5
 * intervals is sorted by starti in ascending order.
 * newInterval.length == 2
 * 0 <= start <= end <= 10^5
 *
 *
 */

// @lc code=start
// intervals already sorted
func merge(intervals [][]int) [][]int {
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

// intervals already sorted
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) != 2 {
		panic("invalid input")
	}
	startNew := newInterval[0]
	// linear search to find position
	i := 0
	for _, v := range intervals {
		start := v[0]
		if start > startNew {
			break
		}
		i++
	}

	tempIntervals := append([][]int{}, intervals[:i]...)
	tempIntervals = append(tempIntervals, newInterval)
	tempIntervals = append(tempIntervals, intervals[i:]...)
	return merge(tempIntervals)
}

// @lc code=end
