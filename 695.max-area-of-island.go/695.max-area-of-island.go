package main

/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 *
 * https://leetcode.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (71.62%)
 * Likes:    8292
 * Dislikes: 183
 * Total Accepted:    648.2K
 * Total Submissions: 904.1K
 * Testcase Example:  '[[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]'
 *
 * You are given an m x n binary matrix grid. An island is a group of 1's
 * (representing land) connected 4-directionally (horizontal or vertical.) You
 * may assume all four edges of the grid are surrounded by water.
 *
 * The area of an island is the number of cells with a value 1 in the island.
 *
 * Return the maximum area of an island in grid. If there is no island, return
 * 0.
 *
 *
 * Example 1:
 *
 *
 * Input: grid =
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 * Output: 6
 * Explanation: The answer is not 11, because the island must be connected
 * 4-directionally.
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[0,0,0,0,0,0,0,0]]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 50
 * grid[i][j] is either 0 or 1.
 *
 *
 */

// @lc code=start
type coord struct {
	x int
	y int
}

func maxAreaOfIsland(grid [][]int) int {

	if len(grid) < 1 {
		return 0
	}

	height := len(grid)
	width := len(grid[0])

	marked := make(map[coord]int)

	var dfs func(c coord)
	curId := 1

	dfs = func(c coord) {
		marked[c] = curId

		x := c.x
		y := c.y

		left := coord{x - 1, y}
		right := coord{x + 1, y}
		bottom := coord{x, y + 1}
		top := coord{x, y - 1}
		neighbors := []coord{}
		for _, v := range []coord{left, right} {
			if v.x >= 0 && v.x < width {
				neighbors = append(neighbors, v)
			}
		}
		for _, v := range []coord{top, bottom} {
			if v.y >= 0 && v.y < height {
				neighbors = append(neighbors, v)

			}
		}
		for _, n := range neighbors {
			if marked[n] < 1 && grid[n.y][n.x] == 1 {
				dfs(n)
			}
		}
	}

	// iterate over all coordinates
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			coord := coord{x, y}
			if marked[coord] < 1 && grid[coord.y][coord.x] == 1 {
				dfs(coord)
				curId++
			}
		}
	}

	islandCounts := make(map[int]int)
	for _, id := range marked {
		if id > 0 {
			islandCounts[id]++
		}
	}

	maxIslands := 0
	for _, ic := range islandCounts  {
		if ic > maxIslands {
			maxIslands = ic
		}
	}
	// now collect the max of islands

	return maxIslands
}

// @lc code=end
