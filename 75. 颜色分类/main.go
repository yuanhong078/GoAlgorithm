package main

import "fmt"

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 2, 1})
}
func sortColors(nums []int) {
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			continue
		} else {
			j := length - 1
			for j != i {
				if nums[i] > nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
				j--
			}
			j = length - 1
			for j != i {
				if nums[i] > nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
				j--
			}
		}
	}

	fmt.Println(nums)

}
