package main

import "fmt"

/*
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X


运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	//board:=[][]byte{{'O','O','O','O','X','X'},{'O','O','O','O','O','O'},{'O','X','O','X','O','O'},{'O','X','O','O','X','O'},{'O','X','O','X','O','O'},{'O','X','O','O','O','O'}}
	for _, v := range board {
		fmt.Println(v)
	}
	fmt.Println()
	solve(board)
	for _, v := range board {
		fmt.Println(v)
	}

}

var dir [][]int

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	dir = [][]int{{0, 1}, {-1, 0}, {1, 0}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j] && inBoard(m, n, i, j) {
				dfs(i, j, board, visited)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j]  {
				board[i][j]='X'
			}
		}
	}
}

func dfs(x, y int, board [][]byte, visited [][]bool)  {
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		newX := x + dir[i][0]
		newY := y + dir[i][1]
		if inArea(newX, newY, len(board), len(board[0])) && board[newX][newY] == 'O' && !visited[newX][newY] {
			 dfs(newX, newY, board, visited)
		}
	}
}

func inArea(x, y int, m, n int) bool {
	if x >= 0 && x < m && y >= 0 && y < n {
		return true
	}
	return false
}

func inBoard(m, n, x, y int) bool {
	if x == 0 || x == m-1 || y == 0 || y == n-1 {
		return true
	}
	return false
}
