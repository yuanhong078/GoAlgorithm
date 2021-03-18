package main

import "fmt"

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var (
	res  [][]int
	used []bool
)

func permute(nums []int) [][]int {
	length := len(nums)
	used = make([]bool, length)
	var p []int
	generatePermutation(nums, 0, p)
	return res
}

func generatePermutation(nums []int, index int, p []int) {
	if len(nums) == index {
		t:=make([]int, len(p))
		copy(t,p)
		res = append(res, t)
		return
	}

	for i := 0; i < len(nums); i++ {

		if !used[i] {
			p = append(p, nums[i])
			used[i] = true
			generatePermutation(nums, index+1, p)
			used[i] = false
			p = p[:len(p)-1]
		}
	}
}

func main() {
	r:=permute([]int{1,1,2})
	fmt.Println(r)
}