package main


import "fmt"

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1:

输入:
[
['1','1','1','1','0'],
['1','1','0','1','0'],
['1','1','0','0','0'],
['0','0','0','0','0']
]
输出: 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

var dir [][]int

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	dir = [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				res++
				fmt.Println(res)
				dfs(i, j, visited, m, n)
			}
		}
	}

	return res
}

func dfs(x, y int, visited [][]bool, m, n int) {
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		newX := x + dir[i][0]
		newY := y + dir[i][1]
		if inArea(newX, newY, m, n) && !visited[newX][newY] {
			dfs(newX, newY, visited, m, n)
		}
	}

}

func inArea(x, y int, m, n int) bool {
	if x >= 0 && x < m && y >= 0 && y < n {
		return true
	}
	return false
}
