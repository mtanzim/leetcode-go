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
	grid   [][]int
	marked [][]bool
}

func traverse(dft *dfsTracker, r, c int) {
	maxRow := len(dft.grid) - 1
	maxCol := len(dft.grid[0]) - 1
	dft.marked[r][c] = true
	neighbors := [][]int{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}}
	isCurRotten := dft.grid[r][c] == rotten
	for _, neighbor := range neighbors {
		nR, nC := neighbor[0], neighbor[1]
		isNeighborEmpty := dft.grid[nR][nC] == empty
		if dft.marked[nR][nC] || nR < 0 || nR > maxRow || nC < 0 || nC > maxCol || isNeighborEmpty {
			continue
		}
		traverse(dft, nR, nC)
		if isCurRotten {
			dft.grid[nR][nC] = rotten
		}
	}
}

func orangesRotting(grid [][]int) int {
	marked := make([][]bool, len(grid))
	for i := range marked {
		marked[i] = make([]bool, len(grid[0]))
	}
	var newGrid [][]int
	copy(newGrid, grid)
	dft := &dfsTracker{newGrid, marked}
	for {
		traverse(dft, 0, 0)
	}
}

// @lc code=end
