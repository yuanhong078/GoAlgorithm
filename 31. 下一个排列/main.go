package main

import (
	"fmt"
	"sort"
)

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nextPermutation([]int{1, 3, 2})
}

//1243 => 1324
func nextPermutation(nums []int) {
	length := len(nums)
	if length == 0 || length == 1 {
		return
	}

	for i := length - 1; i > 0; i-- {
		temp := false
		for j := i - 1; j >= 0; j-- {
			if isMax(nums[j:]) {
				continue
			} else {
				//找出次大的数组
				for k := length - 2; k >= j; k-- {
					for e := k + 1; e < length; e++ {
						if nums[k] < nums[e] {
							nums[k], nums[e] = nums[e], nums[k]
							temp = true
							break
						}
					}
					if temp {
						break
					}
					sort.Ints(nums[k:])
				}
				break
			}
		}

		if temp {
			break
		}
		if i == 1 {
			//变成最小输出
			for start, end := 0, length-1; start < end; start, end = 1+start, end-1 {
				nums[start], nums[end] = nums[end], nums[start]
			}
		}
	}

	fmt.Println(nums)
}

func isMax(num []int) bool {
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] < num[j] {
				return false
			}
		}
	}
	return true
}
