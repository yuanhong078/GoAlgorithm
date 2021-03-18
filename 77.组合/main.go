package main

import "fmt"

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var res [][]int

func main() {
	result := combine(4, 2)
	fmt.Println(result)
}

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return nil
	}

	//res = make([][]int, 0)
	p := make([]int, 0)
	generateCombination(n, k, 1, p)
	return res
}

func generateCombination(n, k, start int, p []int) {

	if len(p) == k {
		e := make([]int, len(p))
		copy(e, p)
		res = append(res, e)
		return
	}

	for i := start; i < n-(k-len(p))+2; i++ {
		p = append(p, i)
		generateCombination(n, k, i+1, p)
		p = p[:len(p)-1]
	}
	return
}
