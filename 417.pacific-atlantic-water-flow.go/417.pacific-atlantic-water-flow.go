package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 *
 * https://leetcode.com/problems/pacific-atlantic-water-flow/description/
 *
 * algorithms
 * Medium (54.56%)
 * Likes:    6522
 * Dislikes: 1251
 * Total Accepted:    367.2K
 * Total Submissions: 673K
 * Testcase Example:  '[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]'
 *
 * There is an m x n rectangular island that borders both the Pacific Ocean and
 * Atlantic Ocean. The Pacific Ocean touches the island's left and top edges,
 * and the Atlantic Ocean touches the island's right and bottom edges.
 *
 * The island is partitioned into a grid of square cells. You are given an m x
 * n integer matrix heights where heights[r][c] represents the height above sea
 * level of the cell at coordinate (r, c).
 *
 * The island receives a lot of rain, and the rain water can flow to
 * neighboring cells directly north, south, east, and west if the neighboring
 * cell's height is less than or equal to the current cell's height. Water can
 * flow from any cell adjacent to an ocean into the ocean.
 *
 * Return a 2D list of grid coordinates result where result[i] = [ri, ci]
 * denotes that rain water can flow from cell (ri, ci) to both the Pacific and
 * Atlantic oceans.
 *
 *
 * Example 1:
 *
 *
 * Input: heights =
 * [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
 * Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
 * Explanation: The following cells can flow to the Pacific and Atlantic
 * oceans, as shown below:
 * [0,4]: [0,4] -> Pacific Ocean
 * [0,4] -> Atlantic Ocean
 * [1,3]: [1,3] -> [0,3] -> Pacific Ocean
 * [1,3] -> [1,4] -> Atlantic Ocean
 * [1,4]: [1,4] -> [1,3] -> [0,3] -> Pacific Ocean
 * [1,4] -> Atlantic Ocean
 * [2,2]: [2,2] -> [1,2] -> [0,2] -> Pacific Ocean
 * [2,2] -> [2,3] -> [2,4] -> Atlantic Ocean
 * [3,0]: [3,0] -> Pacific Ocean
 * [3,0] -> [4,0] -> Atlantic Ocean
 * [3,1]: [3,1] -> [3,0] -> Pacific Ocean
 * [3,1] -> [4,1] -> Atlantic Ocean
 * [4,0]: [4,0] -> Pacific Ocean
 * ⁠      [4,0] -> Atlantic Ocean
 * Note that there are other possible paths for these cells to flow to the
 * Pacific and Atlantic oceans.
 *
 *
 * Example 2:
 *
 *
 * Input: heights = [[1]]
 * Output: [[0,0]]
 * Explanation: The water can flow from the only cell to the Pacific and
 * Atlantic oceans.
 *
 *
 *
 * Constraints:
 *
 *
 * m == heights.length
 * n == heights[r].length
 * 1 <= m, n <= 200
 * 0 <= heights[r][c] <= 10^5
 *
 *
 */

// @lc code=start

func touchesPacific(ri, ci int) bool {
	return ci == 0 || ri == 0
}

func touchesAtlantic(ri, ci, rLen, cLen int) bool {
	return ci == cLen-1 || ri == rLen-1
}

func genKey(ri, ci int) string {
	return fmt.Sprintf("%d|%d", ri, ci)
}

func reverseKey(key string) (int, int) {
	vals := strings.Split(key, "|")
	ri, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	ci, err := strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}
	return ri, ci
}

func dfs(ri int, ci int, heights [][]int, marked map[string]bool) map[string]bool {
	marked[genKey(ri, ci)] = true
	rLen := len(heights)
	cLen := len(heights[0])
	neighbors := [][]int{{ri - 1, ci}, {ri + 1, ci}, {ri, ci - 1}, {ri, ci + 1}}
	for _, neighbor := range neighbors {
		nri, nci := neighbor[0], neighbor[1]
		if !marked[genKey(nri, nci)] && nri >= 0 && nci >= 0 && nri < rLen && nci < cLen && heights[nri][nci] <= heights[ri][ci] {
			// fmt.Println(fmt.Sprintf("curRow %d, curCol %d", nri, nci))
			dfs(nri, nci, heights, marked)
		}
	}
	return marked
}

func pacificAtlantic(heights [][]int) [][]int {

	connectedToOceans := make(map[string]bool)
	rLen := len(heights)
	cLen := len(heights[0])
	for ri, row := range heights {
		for ci := range row {
			markedPath := dfs(ri, ci, heights, make(map[string]bool))

			cPacific := false
			cAtlantic := false
			for key := range markedPath {
				cri, cci := reverseKey(key)
				if touchesAtlantic(cri, cci, rLen, cLen) {
					cAtlantic = cAtlantic || true
				}
				if touchesPacific(cri, cci) {
					cPacific = cPacific || true
				}
			}
			if cPacific && cAtlantic {
				connectedToOceans[genKey(ri, ci)] = true
			}
		}
	}
	res := [][]int{}
	for key := range connectedToOceans {
		ri, ci := reverseKey(key)
		res = append(res, []int{ri, ci})
	}
	return res
}

// @lc code=end
