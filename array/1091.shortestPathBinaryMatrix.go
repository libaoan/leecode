package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 1},
		{1, 0},
	}
	grid = [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	grid = [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 1},
	}
	grid = [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 0},
	}
	grid = [][]int{
		{0, 1, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid))
	for _, gs := range grid {
		fmt.Println(gs)
	}
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid) - 1
	if grid[0][0] != 0 || grid[n][n] != 0 {
		return -1
	}
	return dfs(grid, 0, 0, 0, n)
}

// 全遍历 todo  通过率 56%
func dfs(grid [][]int, i, j, step, n int) int {

	if i == n && j == n {
		return step + 1
	}

	grid[i][j] = 2

	ds := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}

	if i < n && j < n && grid[i+1][j+1] == 0 {
		ds[0] = dfs(grid, i+1, j+1, step, n)
	}

	if i < n && grid[i+1][j] == 0 {
		ds[5] = dfs(grid, i+1, j, step, n)
	}
	if j < n && grid[i][j+1] == 0 {
		ds[7] = dfs(grid, i, j+1, step, n)
	}

	if i > 0 && j < n && grid[i-1][j+1] == 0 {
		ds[1] = dfs(grid, i-1, j+1, step, n)
	}
	if i > 0 && j < n && grid[i-1][j+1] == 0 {
		ds[2] = dfs(grid, i-1, j+1, step, n)
	}
	if i > 0 && j > 0 && grid[i-1][j-1] == 0 {
		ds[3] = dfs(grid, i-1, j-1, step, n)
	}
	if i > 0 && grid[i-1][j] == 0 {
		ds[4] = dfs(grid, i-1, j, step, n)
	}
	if j > 0 && grid[i][j-1] == 0 {
		ds[6] = dfs(grid, i, j-1, step, n)
	}

	best := -1
	fmt.Println(ds, i, j)
	for _, s := range ds {
		if best < 0 {
			best = s
		} else if best < s {
			best = s
		}
	}

	if best < 0 {
		grid[i][j] = 0
	} else {
		best++
	}
	return best
}
