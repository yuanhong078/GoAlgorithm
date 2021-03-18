package main

import "fmt"

/*
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

    所有数字都是正整数。
    解集不能包含重复的组合。

示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]

示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var res [][]int

func main() {
	r:=combinationSum3(3,9)
	fmt.Println(r)
}

func combinationSum3(k int, n int) [][]int {
	if k > 9 || n < 1 {
		return nil
	}
	p := make([]int, 0)
	gengerateCombinationSum3(k, n, 1, p)
	return res
}

func gengerateCombinationSum3(k, n, start int, p []int) {

	if n == 0 && len(p) == k {
		e := make([]int, k)
		copy(e, p)
		res = append(res, e)
		return
	}

	if len(p) >= k && n != 0 {
		return
	}

	for i := start; i <= 9; i++ {
		if n >= i {
			p = append(p, i)
			gengerateCombinationSum3(k, n-i, i+1, p)
			p = p[:len(p)-1]
		} else {
			break
		}

	}
}
