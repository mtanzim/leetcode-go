package main

/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 *
 * https://leetcode.com/problems/flood-fill/description/
 *
 * algorithms
 * Easy (60.21%)
 * Likes:    6032
 * Dislikes: 580
 * Total Accepted:    607.8K
 * Total Submissions: 1M
 * Testcase Example:  '[[1,1,1],[1,1,0],[1,0,1]]\n1\n1\n2'
 *
 * An image is represented by an m x n integer grid image where image[i][j]
 * represents the pixel value of the image.
 *
 * You are also given three integers sr, sc, and color. You should perform a
 * flood fill on the image starting from the pixel image[sr][sc].
 *
 * To perform a flood fill, consider the starting pixel, plus any pixels
 * connected 4-directionally to the starting pixel of the same color as the
 * starting pixel, plus any pixels connected 4-directionally to those pixels
 * (also with the same color), and so on. Replace the color of all of the
 * aforementioned pixels with color.
 *
 * Return the modified image after performing the flood fill.
 *
 *
 * Example 1:
 *
 *
 * Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
 * Output: [[2,2,2],[2,2,0],[2,0,1]]
 * Explanation: From the center of the image with position (sr, sc) = (1, 1)
 * (i.e., the red pixel), all pixels connected by a path of the same color as
 * the starting pixel (i.e., the blue pixels) are colored with the new color.
 * Note the bottom corner is not colored 2, because it is not 4-directionally
 * connected to the starting pixel.
 *
 *
 * Example 2:
 *
 *
 * Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
 * Output: [[0,0,0],[0,0,0]]
 * Explanation: The starting pixel is already colored 0, so no changes are made
 * to the image.
 *
 *
 *
 * Constraints:
 *
 *
 * m == image.length
 * n == image[i].length
 * 1 <= m, n <= 50
 * 0 <= image[i][j], color < 2^16
 * 0 <= sr < m
 * 0 <= sc < n
 *
 *
 */

// @lc code=start
type coord struct {
	x int
	y int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	height := len(image)
	width := len(image[0])

	source := coord{sc, sr}
	marked := make(map[coord]bool)
	var dfs func(v coord)
	dfs = func(v coord) {
		marked[v] = true
		neighbors := []coord{}
		left := coord{v.x - 1, v.y}
		right := coord{v.x + 1, v.y}
		top := coord{v.x, v.y - 1}
		bottom := coord{v.x, v.y + 1}

		for _, c := range []coord{left, right} {
			if c.x > -1 && c.x < width && image[c.y][c.x] == image[v.y][v.x] && !marked[c] {
				neighbors = append(neighbors, c)
			}
		}
		for _, c := range []coord{top, bottom} {
			if c.y > -1 && c.y < height && image[c.y][c.x] == image[v.y][v.x] && !marked[c] {
				neighbors = append(neighbors, c)
			}
		}

		for _, c := range neighbors {
			dfs(c)
		}

	}

	dfs(source)

	transformed := make([][]int, height)
	for y := range image {
		transformed[y] = make([]int, width)
		for x := range image[y] {
			if marked[coord{x, y}] {
				transformed[y][x] = color
			} else {
				transformed[y][x] = image[y][x]

			}
		}
	}

	return transformed

}

// @lc code=end
