package main

import (
	"fmt"
	"sort"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：


	所有数字（包括 target）都是正整数。
	解集不能包含重复的组合。


示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var res [][]int

func main() {
	r:=combinationSum([]int{2,3,5},8)
	fmt.Println(r)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	if target<candidates[0]{
		return nil
	}
	var p []int
	generateCombinationSum(candidates, target,0,p)
	return res
}

func generateCombinationSum(candidates []int, target int, start int, p []int) {
	if target == 0 {
		e := make([]int, len(p))
		copy(e, p)
		res = append(res, e)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			return
		}
		p = append(p, candidates[i])
		generateCombinationSum(candidates, target-candidates[i], i, p)
		p = p[:len(p)-1]
	}
}
