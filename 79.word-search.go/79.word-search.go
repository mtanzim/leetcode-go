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

func dfs(board [][]byte, word string, marked map[coord]bool, curCoord coord) bool {

	if len(board) < 1 {
		return true
	}
	if len(word) < 1 {
		return true
	}

	x := curCoord.x
	y := curCoord.y
	first := string(word[0])
	curLetter := string(board[y][x])

	if curLetter != first {
		return false
	}

	if curLetter == first && len(word) == 1 {
		return true
	}
	marked[curCoord] = true

	height := len(board)
	width := len(board[0])

	left := coord{x - 1, y}
	right := coord{x + 1, y}
	bottom := coord{x, y + 1}
	top := coord{x, y - 1}
	neighbors := make(map[coord]bool)
	for _, v := range []coord{left, right} {
		if v.x >= 0 && v.x < width && !marked[v] {
			neighbors[v] = true
		}
	}
	for _, v := range []coord{top, bottom} {
		if v.y >= 0 && v.y < height && !marked[v] {
			neighbors[v] = true
		}
	}
	found := false
	for c := range neighbors {
		markedCopy := make(map[coord]bool)
		for k, v := range marked {
			markedCopy[k] = v
		}
		found = found || dfs(board, word[1:], markedCopy, c)
	}
	return found
}

func exist(board [][]byte, word string) bool {
	height := len(board)
	width := len(board[0])

	doesExist := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			doesExist = doesExist || dfs(board, word, make(map[coord]bool), coord{x, y})
		}
	}

	return doesExist

}

// @lc code=end
