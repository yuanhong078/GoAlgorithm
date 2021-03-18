package main

import "fmt"

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]


示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]



说明：


	输出结果中的每个元素一定是唯一的。
	我们可以不考虑输出结果的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	result:=intersection(nums1, nums2)
	fmt.Println(result)
}

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; ok {
			continue
		}
		set[v] = true
	}
	result := []int{}
	for _, v := range nums2 {
		if t, ok := set[v]; ok {
			if t == true {
				result = append(result,v)
				set[v]=false
				continue
			}
		}
	}

	return result
}
