package main

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

*/

func main() {

}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[m-1])
	for x := m - 1; x >= 0; x-- {
		for y := n - 1; y >= 0; y-- {
			x1, y1 := x+1, y
			x2, y2 := x, y+1

			if inArea(x1, y1, m, n) && inArea(x2, y2, m, n) {
				grid[x][y] = grid[x][y] + min(grid[x1][y1], grid[x2][y2])
			} else { //在边沿
				if inArea(x1, y1, m, n) {
					grid[x][y] = grid[x][y] + grid[x1][y1]
				} else if inArea(x2, y2, m, n) {
					grid[x][y] = grid[x][y] + grid[x2][y2]
				}
			}
		}
	}

	return grid[0][0]
}

func inArea(x, y, m, n int) bool {
	if x > m-1 || y > n-1 {
		return false
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
