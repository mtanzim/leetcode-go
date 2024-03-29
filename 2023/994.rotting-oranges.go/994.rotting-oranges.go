package main

/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 *
 * https://leetcode.com/problems/rotting-oranges/description/
 *
 * algorithms
 * Medium (53.13%)
 * Likes:    10916
 * Dislikes: 349
 * Total Accepted:    629.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[2,1,1],[1,1,0],[0,1,1]]'
 *
 * You are given an m x n grid where each cell can have one of three
 * values:
 *
 *
 * 0 representing an empty cell,
 * 1 representing a fresh orange, or
 * 2 representing a rotten orange.
 *
 *
 * Every minute, any fresh orange that is 4-directionally adjacent to a rotten
 * orange becomes rotten.
 *
 * Return the minimum number of minutes that must elapse until no cell has a
 * fresh orange. If this is impossible, return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
 * Output: -1
 * Explanation: The orange in the bottom left corner (row 2, column 0) is never
 * rotten, because rotting only happens 4-directionally.
 *
 *
 * Example 3:
 *
 *
 * Input: grid = [[0,2]]
 * Output: 0
 * Explanation: Since there are already no fresh oranges at minute 0, the
 * answer is just 0.
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10
 * grid[i][j] is 0, 1, or 2.
 *
 *
 */

// @lc code=start

const (
	empty  = 0
	fresh  = 1
	rotten = 2
)

type dfsTracker struct {
	grid    [][]int
	marked  [][]bool
	willRot [][]int
}

func traverse(dft *dfsTracker, r, c int) {
	if dft.marked[r][c] {
		return
	}
	maxRow := len(dft.grid) - 1
	maxCol := len(dft.grid[0]) - 1
	dft.marked[r][c] = true
	neighbors := [][]int{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}}
	filteredNeighbors := [][]int{}
	for _, neighbor := range neighbors {
		nR, nC := neighbor[0], neighbor[1]
		if nR < 0 || nR > maxRow || nC < 0 || nC > maxCol || dft.marked[nR][nC] || dft.grid[nR][nC] == empty {
			continue
		}
		filteredNeighbors = append(filteredNeighbors, []int{nR, nC})
	}
	// check who will rot
	if dft.grid[r][c] == rotten {
		for _, neighbor := range neighbors {
			nR, nC := neighbor[0], neighbor[1]
			if nR < 0 || nR > maxRow || nC < 0 || nC > maxCol || dft.grid[nR][nC] == empty {
				continue
			}
			dft.willRot = append(dft.willRot, neighbor)
		}
	}
	for _, neighbor := range filteredNeighbors {
		nR, nC := neighbor[0], neighbor[1]
		traverse(dft, nR, nC)
	}

}

func orangesRotting(grid [][]int) int {

	dft := &dfsTracker{copyGrid(grid), makeMarked(grid), [][]int{}}

	minutes := 0

	if countFresh(grid) == 0 {
		return minutes
	}

	for {
		prevFresh := countFresh(dft.grid)
		for ri, row := range dft.grid {
			for ci := range row {
				traverse(dft, ri, ci)
			}
		}
		minutes++
		for _, coord := range dft.willRot {
			rotR, rotC := coord[0], coord[1]
			dft.grid[rotR][rotC] = rotten
		}
		newFresh := countFresh(dft.grid)
		if newFresh == 0 {
			return minutes
		}
		if newFresh == prevFresh {
			return -1
		}

		dft.marked = makeMarked(grid)
		dft.willRot = [][]int{}
	}
}

func copyGrid(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for ri, rv := range grid {
		newGrid[ri] = make([]int, len(rv))
		for ci := range rv {
			newGrid[ri][ci] = grid[ri][ci]
		}
	}
	return newGrid
}

func makeMarked(grid [][]int) [][]bool {
	newMarked := make([][]bool, len(grid))
	for i := range grid {
		newMarked[i] = make([]bool, len(grid[0]))
	}
	return newMarked
}

func countFresh(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v == fresh {
				count++
			}
		}
	}
	return count
}

// @lc code=end
