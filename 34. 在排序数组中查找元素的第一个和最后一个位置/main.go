package main

import (
	"fmt"
	"sort"
)

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]

示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 4, 4
//1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
//3

func main() {
	res := searchRange([]int{0, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 4, 4}, 3)
	fmt.Println(res)
}

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}

	res := make([]int, 0)

	res=twosearch(nums, target, 0)
	sort.Ints(res)
	fmt.Println("---", res)

	return []int{res[0], res[len(res)-1]}
}

func twosearch(nums []int, target, power int) (res []int) {
	res = make([]int, 0)
	length := len(nums)
	if length == 1 {
		if nums[0] != target {
			res = append(res, []int{-1, -1}...)
		} else {
			res = append(res, 0+power)
		}
		return
	}
	mid := length / 2
	if nums[mid] > target {
		res = twosearch(nums[:mid], target, power)
	} else if nums[mid] < target {
		res = twosearch(nums[mid:], target, mid+power)
	} else if nums[mid] == target {
		fmt.Println("power:", power)
		for i := mid; i > -1; i-- {
			if nums[i] != target {
				break
			}
			res = append(res, i+power)
			fmt.Println("res:", res)
		}
		for i := mid + 1; i < length; i++ {
			if nums[i] != target {
				break
			}
			res = append(res, i+power)
		}
	}
	return
}
