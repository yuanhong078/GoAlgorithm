package main

import (
	"fmt"
	"sort"
)

/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var res [][]int

func main() {
	r:=subsetsWithDup([]int{1,2,2})
	fmt.Println(r)
}

func subsetsWithDup(nums []int) [][]int {

	res = make([][]int, 0)
	if len(nums) == 0 {
		res = append(res, []int{})
		return res
	}
	res = append(res, []int{})
	used := make([]bool, len(nums))
	p := make([]int, 0)
	sort.Ints(nums)
	generateSubsetsWithDup(nums, p, 0, used)

	return res
}

func generateSubsetsWithDup(nums, p []int, start int, used []bool) {

	for i := start; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		p = append(p, nums[i])
		e := make([]int, len(p))
		copy(e, p)
		res = append(res, e)
		used[i] = true
		generateSubsetsWithDup(nums, p, i+1, used)
		p = p[:len(p)-1]
		used[i] = false
	}
}

