package main

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。

 

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

/*
     [-1],
    [2,3],
   [1,-1,-3],
*/

//错误解法
func minimumTotal9(triangle [][]int) int {
	k := 0

	sum := triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		if triangle[i][k] < triangle[i][k+1] {
			sum += triangle[i][k]
		} else {
			k++
			sum += triangle[i][k]
		}
	}
	return sum
}

//自底向上，计算每个[][]的最小值，直到[0][0]
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
