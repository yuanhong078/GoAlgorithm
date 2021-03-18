package main

/*
给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。

规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。

请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。
提示：
	输出坐标的顺序不重要
	m 和 n 都小于150
示例：
给定下面的 5x5 矩阵:

  太平洋 ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * 大西洋

返回:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pacific-atlantic-water-flow
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

var dir, res [][]int
var visited [][]bool

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	m := len(matrix)
	n := len(matrix[0])
	dir = [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pm := make(map[rune]bool, 0)
			pm['p'] = false
			pm['a'] = false
			dfs(i, j, matrix, pm)
			visited[i][j]=false
			if pm['p'] == true && pm['a'] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(x, y int, matrix [][]int, pm map[rune]bool) {
	visited[x][y]=true
	m := len(matrix)
	n := len(matrix[0])
	if isPacific(x, y) {
		pm['p'] = true
	}
	if isAtlantic(x, y, m, n) {
		pm['a'] = true
	}
	if pm['p'] == true && pm['a'] {
		return
	}

	for i := 0; i < 4; i++ {
		newX := x + dir[i][0]
		newY := y + dir[i][1]
		if inArea(newX, newY, m, n) && !visited[newX][newY] && matrix[x][y] >= matrix[newX][newY] {
			dfs(newX, newY, matrix, pm)
			visited[newX][newY] = false
		}
	}

}

func inArea(x, y int, m, n int) bool {
	if x >= 0 && y >= 0 && x < m && y < n {
		return true
	}
	return false
}

func isPacific(x, y int) bool {
	if y == 0 || x == 0 {
		return true
	}
	return false
}

func isAtlantic(x, y int, m, n int) bool {
	if y == n-1 || x == m-1 {
		return true
	}
	return false
}
