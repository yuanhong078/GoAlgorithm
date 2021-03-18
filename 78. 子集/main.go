package main

import "fmt"

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var res [][]int

func main() {
	r:=subsets([]int{1,2,3})
	fmt.Println(r)
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		res = append(res, []int{})
		return res
	}else{
		res = append(res, []int{})
	}
	res=make([][]int,0)
	p := make([]int, 0)
	generateSubsets(nums, p, 0)
	return res
}

func generateSubsets(nums, p []int, start int) {

	for i := start; i < len(nums); i++ {
		p = append(p, nums[i])
		e := make([]int, len(p))
		copy(e, p)
		res = append(res, e)
		generateSubsets(nums, p, i+1)
		p = p[:len(p)-1]
	}
}
