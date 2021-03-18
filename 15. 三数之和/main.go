package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for k, v := range res {
		fmt.Println(k, "   :", v)
	}
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	fmt.Println(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}

	exists := make(map[int][]int)
	var res [][]int
	var items []int

	for k, v := range nums {
		if v > 0 {
			break
		}
		if isContain(items, v) {
			continue
		}
		items = append(items, v) //将该值加入已经遍历过的数组，不再遍历
		sub := 0 - v
		for i := k + 1; i < len(nums)-1; i++ {
			a := nums[i]
			b := sub - a

			if isContain(nums[i+1:], b) {

				if exists[a] == nil {
					exists[a] = make([]int, 0)
				} else if isContain(exists[a], b) {
					continue
				}
				exists[a] = append(exists[a], b, v)
				exists[b] = append(exists[b], a, v)
				res = append(res, []int{v, a, b})
			}
		}
	}
	return res
}

func isContain(nums []int, v int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == v {
			return true
		}
	}
	return false
}

//func sort(nums []int) []int {
//
//	for i := 0; i < len(nums)-1; i++ {
//		min := nums[i]
//		temp := i
//		for j := i + 1; j < len(nums); j++ {
//			if nums[j] < min {
//				min = nums[j]
//				temp = j
//			}
//		}
//		if temp != i {
//			nums[i], nums[temp] = nums[temp], nums[i]
//		}
//	}
//	return nums
//}
