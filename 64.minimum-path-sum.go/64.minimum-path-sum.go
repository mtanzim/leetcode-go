package main

import "fmt"

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (60.51%)
 * Likes:    8777
 * Dislikes: 113
 * Total Accepted:    799K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right, which minimizes the sum of all numbers along its
 * path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
 * Output: 7
 * Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[1,2,3],[4,5,6]]
 * Output: 12
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 100
 *
 *
 */

type weight = int

// vertex indexed adjacency list
type graph struct {
	v   int
	e   int
	adj []map[int]weight
}

func (g *graph) addEdge(from, to, weight int) {
	g.adj[from][to] = weight
}

func newGraph(v int) *graph {

	g := &graph{v: v, e: 0}
	g.adj = make([]map[int]int, v)
	for i := 0; i < v; i++ {
		g.adj[i] = make(map[int]int)
	}
	return g
}

type coord struct {
	x int
	y int
}

func dfsOrder(g *graph) []int {
	marked := make([]bool, g.v)
	edgeTo := make([]int, g.v)
	topoStack := []int{}
	for i := range edgeTo {
		edgeTo[i] = -1
	}
	var dfs func(v int)
	dfs = func(v int) {
		marked[v] = true
		adj := g.adj[v]
		neighbors := make([]int, len(adj))
		i := 0
		for k := range adj {
			neighbors[i] = k
			i++
		}
		for _, neighbor := range neighbors {
			if !marked[neighbor] {
				edgeTo[neighbor] = v
				dfs(neighbor)
			}
		}
		topoStack = append(topoStack, v)
	}
	dfs(0)
	return topoStack

}

// @lc code=start
func minPathSum(grid [][]int) int {
	idHM := make(map[coord]int)
	id := 0
	for y := range grid {
		for x := range grid[y] {
			idHM[coord{x, y}] = id
			id++
		}
	}
	fmt.Printf("coord map: %v\n", idHM)

	g := newGraph(len(idHM))
	for y := range grid {
		for x := range grid[y] {
			from := idHM[coord{x, y}]
			right, rightOk := idHM[coord{x + 1, y}]
			down, downOk := idHM[coord{x, y + 1}]

			if rightOk {
				rightWeight := grid[y][x+1]
				g.addEdge(from, right, rightWeight)
			}
			if downOk {
				downWeight := grid[y+1][x]
				g.addEdge(from, down, downWeight)
			}
		}
	}

	topoOrder := dfsOrder(g)
	fmt.Printf("topo stack: %v\n", topoOrder)

	return 0
}

// @lc code=end
