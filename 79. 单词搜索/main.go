package main

import "fmt"

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true

链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var dist = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

var visited [][]bool

//var board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
var board = [][]byte{{'a', 'a'}}

func main() {
	fmt.Println(exist(board, "aaa"))
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		visited[i]=make([]bool,n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchWord(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte, word string, x, y int, target int) bool {

	if target == len(word)-1 {
		//if string(word[target]) == string(board[x][y]) {
		//	return true
		//}
		//return false
		return string(word[target]) == string(board[x][y])
	}

	if string(word[target]) == string(board[x][y]) {

		visited[x][y] = true
		for i := 0; i < len(dist); i++ {
			newX := x + dist[i][0]
			newY := y + dist[i][1]
			if inArea(newX, newY, board) && !visited[newX][newY] {
				if searchWord(board, word, newX, newY, target+1) {
					return true
				}
			}
		}
		visited[x][y] = false
	}
	return false
}

func inArea(x, y int, board [][]byte) bool {
	if x >= 0 && x < len(board) && y < len(board[0]) && y >= 0 {
		return true
	}
	return false
}
