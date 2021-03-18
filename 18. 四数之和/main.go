package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	res := fourSum([]int{1,-2,-5,-4,-3,3,3,5}, -11)
	fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	fmt.Println(nums)
	if  4*nums[len(nums)-1] < target || len(nums) < 4 {
		return nil
	}

	res := make([][]int, 0)

	result := subSum(0, 3, target, nums)
	for _, v := range result {
		res = append(res, v...)
	}

	return res
}

func subSum(num int, list, target int, nums []int) (map[int][][]int) {

	res := make(map[int][][]int)

	if list != 0 {

		for k, value := range nums {
			if k > 0 && value == nums[k-1] {
				continue
			}
			if k+list == len(nums) {
				break
			}

			numList := subSum(k+1, list-1, target-value, nums[k+1:])

			if numList != nil { //如果找到合适的数组
				for _, arr := range numList[k] {
					arr = append(arr, value) //维护返回的数组集合
					res[num-1] = append(res[num-1], arr)
				}
			}
		}

		return res
	}

	i := make([]int, 0)
	for _, value := range nums {
		if value == target {
			i = append(i, value)
			res[num-1] = append(res[num-1], i)
			return res
		}
	}
	return nil

}
