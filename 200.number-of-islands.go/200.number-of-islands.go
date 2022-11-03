package main

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (56.06%)
 * Likes:    17702
 * Dislikes: 406
 * Total Accepted:    1.9M
 * Total Submissions: 3.5M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given an m x n 2D binary grid grid which represents a map of '1's (land) and
 * '0's (water), return the number of islands.
 *
 * An island is surrounded by water and is formed by connecting adjacent lands
 * horizontally or vertically. You may assume all four edges of the grid are
 * all surrounded by water.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 300
 * grid[i][j] is '0' or '1'.
 *
 *
 */

// @lc code=start
type coord struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {

	if len(grid) < 1 {
		return 0
	}

	height := len(grid)
	width := len(grid[0])

	marked := make(map[coord]bool)

	var dfs func(c coord)

	dfs = func(c coord) {
		marked[c] = true

		x := c.x
		y := c.y

		left := coord{x - 1, y}
		right := coord{x + 1, y}
		bottom := coord{x, y + 1}
		top := coord{x, y - 1}
		neighbors := []coord{}
		for _, v := range []coord{left, right} {
			if v.x > 0 && v.x < width {
				neighbors = append(neighbors, v)
			}
		}
		for _, v := range []coord{top, bottom} {
			if v.y > 0 && v.y < height {
				neighbors = append(neighbors, v)

			}
		}
		for _, n := range neighbors {
			if !marked[n] && grid[n.y][n.x] != byte('0') {
				dfs(n)
			}
		}
	}

	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			coord := coord{x, y}
			// fmt.Printf("%v: %c", coord, grid[y][x])
			if !marked[coord] && grid[coord.y][coord.x] == byte('1') {
				dfs(coord)
				count++

			}
		}
	}
	return count
}

// @lc code=end
