package main

import "fmt"

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 *
 * https://leetcode.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (45.33%)
 * Likes:    11990
 * Dislikes: 469
 * Total Accepted:    1.1M
 * Total Submissions: 2.4M
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of numCourses courses you have to take, labeled from 0 to
 * numCourses - 1. You are given an array prerequisites where prerequisites[i]
 * = [ai, bi] indicates that you must take course bi first if you want to take
 * course ai.
 *
 *
 * For example, the pair [0, 1], indicates that to take course 0 you have to
 * first take course 1.
 *
 *
 * Return true if you can finish all courses. Otherwise, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0. So it is possible.
 *
 *
 * Example 2:
 *
 *
 * Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0, and to take course 0 you
 * should also have finished course 1. So it is impossible.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * All the pairs prerequisites[i] are unique.
 *
 *
 */

// @lc code=start

type Adj = [][]int

func dfs(adj Adj, marked map[int]bool, cycleTracker *cycleTracker, v int) bool {
	marked[v] = true
	cycleTracker.onStack[v] = true
	neighbors := adj[v]
	foundCycle := false
	for _, neighbor := range neighbors {
		if !marked[neighbor] {
			curFoundCycle := cycleTracker.onStack[neighbor]
			foundCycle = curFoundCycle || dfs(adj, marked, cycleTracker, neighbor)
		} else if cycleTracker.onStack[neighbor] {
			return true
		}
	}
	cycleTracker.onStack[v] = false
	return foundCycle
}

type cycleTracker struct {
	onStack map[int]bool
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(Adj, numCourses)
	for _, prereq := range prerequisites {
		to, from := prereq[0], prereq[1]
		adj[from] = append(adj[from], to)
	}
	foundCycle := false
	for vertex := range adj {
		tracker := &cycleTracker{make(map[int]bool)}
		hasCycle := dfs(adj, make(map[int]bool), tracker, vertex)
		foundCycle = hasCycle
		if foundCycle {
			break
		}
		fmt.Printf("vertex %d cycle: %v\n", vertex, hasCycle)
	}
	return !foundCycle
}

// @lc code=end
