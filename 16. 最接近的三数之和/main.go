package main

import (
	"fmt"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，
使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。



示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
}

func threeSumClosest(nums []int, target int) int {

	min := nums[1] + nums[len(nums)-1] + nums[0] - target
	res := nums[1] + nums[len(nums)-1] + nums[0]
	sort.Ints(nums)
	for s := 0; s < len(nums)-2; s++ {
		if s != 0 && nums[s-1] == nums[s] {
			continue
		}
		sub := target - nums[s]
		for i, j := s+1, len(nums)-1; i < j; { //-4,-1,1,2

			dis := sub - nums[i] - nums[j]
			fmt.Println("nums:", nums[s], "  ", nums[i], "   ", nums[j], "   ", dis)

			if dis == 0 {
				return target
			}
			if dis*dis < min*min || dis*dis == min*min { //取绝对值
				min = dis
				res = nums[s] + nums[i] + nums[j]
			}
			if dis > 0 { //sum<target
				i++
				continue
			} else {
				j--
				continue
			}

		}
	}
	return res
}
