package main

import (
	"fmt"
	"time"
)

//给定一个整数数组 nums 和一个目标值 target，
// 请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
/*给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	startTime := time.Now().UnixNano()
	res := twoSum2([]int{1, 1, 7, 0, 3}, 8)
	fmt.Println(res)
	endTime := time.Now().UnixNano()
	speed := endTime - startTime
	fmt.Println(speed)
}

//遍历两边
func twoSum(nums []int, target int) []int {

	for k, v := range nums {
		arr := nums[k+1:]
		if len(arr) == 0 {
			break
		}
		for k2, v2 := range arr {
			if v+v2 == target {
				res := []int{k, k2 + k + 1}
				return res
			}
		}
	}
	return nil
}

//遍历一遍
func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		complement := target - v
		if value, ok := numMap[complement]; !ok {
			numMap[v] = k
		} else {
			res := []int{k, value}
			return res
		}

	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		complement := target - v
		value, ok := numMap[complement]
		if ok {
			res := []int{k, value}
			return res
		}
		numMap[v] = k
	}
	return nil
}
