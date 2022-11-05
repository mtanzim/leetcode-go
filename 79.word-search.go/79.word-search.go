package main

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (39.86%)
 * Likes:    11432
 * Dislikes: 450
 * Total Accepted:    1.1M
 * Total Submissions: 2.8M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given an m x n grid of characters board and a string word, return true if
 * word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cells,
 * where adjacent cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once.
 *
 *
 * Example 1:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCCED"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "SEE"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCB"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n = board[i].length
 * 1 <= m, n <= 6
 * 1 <= word.length <= 15
 * board and word consists of only lowercase and uppercase English letters.
 *
 *
 *
 * Follow up: Could you use search pruning to make your solution faster with a
 * larger board?
 *
 */

// @lc code=start

type coord struct {
	x int
	y int
}

func exist(board [][]byte, word string) bool {
	if len(board) < 1 {
		return true
	}
	if len(word) < 1 {
		return true
	}

	height := len(board)
	width := len(board[0])

	if len(word) == 1 {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if rune(board[y][x]) == rune(word[0]) {
					return true
				}
			}
		}
	}
	return false

	// var dfs func(c coord, marked map[coord]bool, curWord []rune)

	// dfs = func(c coord, marked map[coord]bool, curWord []rune) {
	// 	marked[c] = true
	// 	letter := board[c.y][c.x]
	// 	newWord := append([]rune{}, curWord...)
	// 	newWord = append(newWord, rune(letter))
	// 	words = append(words, string(newWord))

	// 	x := c.x
	// 	y := c.y

	// 	left := coord{x - 1, y}
	// 	right := coord{x + 1, y}
	// 	bottom := coord{x, y + 1}
	// 	top := coord{x, y - 1}
	// 	neighbors := []coord{}
	// 	for _, v := range []coord{left, right} {
	// 		if v.x >= 0 && v.x < width {
	// 			neighbors = append(neighbors, v)
	// 		}
	// 	}
	// 	for _, v := range []coord{top, bottom} {
	// 		if v.y >= 0 && v.y < height {
	// 			neighbors = append(neighbors, v)

	// 		}
	// 	}
	// 	for _, n := range neighbors {
	// 		if !marked[n] && board[n.y][n.x] == byte('1') {
	// 			dfs(n, marked, newWord)
	// 		}
	// 	}
	// }

	// marked := make(map[coord]bool)

	// first := rune(word[0])
	// words := []string{}

	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		coord := coord{x, y}
	// 		if !marked[coord] && board[coord.y][coord.x] == byte('1') {
	// 			dfs(coord)

	// 		}
	// 	}
	// }
	// return false

}

// @lc code=end
