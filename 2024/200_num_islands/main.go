package main

import (
	"slices"
)

type Coord struct {
	ri int
	ci int
}

func dfs(grid [][]byte, visited map[Coord]struct{}, curCoord Coord) {
	if _, ok := visited[curCoord]; ok {
		return
	}
	rl := len(grid)
	cl := len(grid[0])
	visited[curCoord] = struct{}{}
	allNeighbors := []Coord{
		{ri: curCoord.ri + 1, ci: curCoord.ci},
		{ri: curCoord.ri - 1, ci: curCoord.ci},
		{ri: curCoord.ri, ci: curCoord.ci + 1},
		{ri: curCoord.ri, ci: curCoord.ci - 1},
	}
	validNs := slices.DeleteFunc(allNeighbors, func(cn Coord) bool {
		if cn.ci < 0 || cn.ci >= cl {
			return true
		}
		if cn.ri < 0 || cn.ri >= rl {
			return true
		}
		if _, ok := visited[cn]; ok {
			return true
		}
		if v := grid[cn.ri][cn.ci]; v == '0' {
			return true
		}
		return false

	})
	for _, v := range validNs {
		dfs(grid, visited, v)
	}

}

func numIslands(grid [][]byte) int {
	visited := make(map[Coord]struct{})
	count := 0
	for ri, row := range grid {
		for ci, v := range row {
			cc := Coord{ci: ci, ri: ri}

			_, isVisited := visited[cc]
			if v == '1' && !isVisited {
				dfs(grid, visited, cc)
				count++
			}
		}
	}
	return count
}
