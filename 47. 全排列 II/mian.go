package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var (
	res  [][]int
	used []bool
)

func main() {
	r := permuteUnique([]int{1,1,2,1})
	fmt.Println(r)
}

func permuteUnique(nums []int) [][]int {
	length := len(nums)
	used = make([]bool, length)
	var p []int
	sort.Ints(nums)
	generatePermutation(nums, 0, p)
	return res
}

func generatePermutation(nums []int, index int, p []int) {
	if len(nums) == index {
		t := make([]int, len(p))
		copy(t, p)
		res = append(res, t)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1]&&!used[i-1] {
			continue
		}

		if !used[i] {
			p = append(p, nums[i])
			used[i] = true
			generatePermutation(nums, index+1, p)
			used[i] = false
			p = p[:len(p)-1]
		}
	}
}
