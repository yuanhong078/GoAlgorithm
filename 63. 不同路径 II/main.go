package main

import "fmt"

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if m == 1 || n == 1 {
		return 1
	}
	path = make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}

	path[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = findPath(i, j, m, n,obstacleGrid)
		}
	}

	return path[m-1][n-1]
}

var path [][]int

func findPath(x, y int, m, n int, obstacleGrid [][]int) int {

	if x == 0 && y == 0 {
		return 1
	}
	r := 0
	rx, ry := x, y-1
	dx, dy := x-1, y
	if isArea(rx, ry, m, n, obstacleGrid) {
		fmt.Println("r:", r, "(", rx, ry)
		fmt.Println(len(path[rx]))
		r = path[rx][ry]

	}

	if isArea(dx, dy, m, n, obstacleGrid) {
		fmt.Println("d:", r, "(", dx, dy)
		r += path[dx][dy]

	}
	return r
}

func isArea(x, y int, m, n int, obstacleGrid [][]int) bool {
	if x <= m && y <= n && x >= 0 && y >= 0 && obstacleGrid[x][y] == 0 {
		return true
	}
	return false
}
