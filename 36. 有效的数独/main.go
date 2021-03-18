package main

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。


	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	board := [][]string{{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"}}

	board1 := make([][]byte, 0)

	for k, v := range board {
		for k2, v2 := range v {
			v3 := []byte(v2)
			board1[k][k2] = v3
		}

	}

	result := isValidSudoku()
}

func isValidSudoku(board [][]byte) bool {

	//判断列[0,0],[1,0]
	for k1, v1 := range board {
		sodu := make(map[byte]int, 0)
		for i := 0; i < len(v1); i++ {
			v := board[i][k1]
			if string(v) != "." {
				if _, ok := sodu[v]; ok {
					return false
				}
			}
			sodu[v] += 1
		}
	}
	//判断行
	for _, v1 := range board {
		sodu := make(map[byte]int, 0)
		for _, v2 := range v1 {
			if string(v2) != "." {
				if _, ok := sodu[v2]; ok {
					return false
				}
			}

			sodu[v2] += 1
		}
	}
	//判断三角
	for k2 := 0; k2 < len(board)-1; k2 += 3 {
		k := 3
		for i := 0; i < len(board)-1; i += k {
			k1 := i
			sodu := make(map[byte]int, 0)
			for k1 < i+k {
				for j := k2; j < j+k; j += 1 {
					v := board[k1][j]
					if string(v) != "." {
						if _, ok := sodu[v]; ok {
							return false
						}
					}
					sodu[v] += 1
				}
			}
			k1++
		}
	}
	return true
}
