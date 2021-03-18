package main

import "fmt"

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	len:=removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(len)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	for i := 1; i < len(nums); {
		//fmt.Println(nums)
		//fmt.Println("i-1:",nums[i-1]," i:",nums[i])
		if nums[i-1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}

	for k, v := range nums {
		fmt.Println("k:", k, " v:", v)
	}
	return len(nums)
}
