package main

import "fmt"

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

说明:
	初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)

}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	k := 0
	sum := make([]int, 0)
	i := 0
	for ; i < n && k < m;
	{
		if nums2[i] > nums1[k] {
			sum = append(sum, nums1[k])
			k++
			continue
		} else if nums2[i] < nums1[k] {
			sum = append(sum, nums2[i])
			i++
			continue
		} else {
			sum = append(sum, nums1[k], nums2[i])
			i++
			k++
			continue
		}
	}
	if k == m {
		sum = append(sum, nums2[i:n]...)
	} else if i == n {
		sum = append(sum, nums1[k:m]...)
	}
	for k, v := range sum {
		nums1[k] = v
	}
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m - 1
	i := n - 1
	end := m + n - 1
	for ; i >= 0 && k >= 0;
	{
		if nums2[i] > nums1[k] {
			nums1[end] = nums2[i]
			end--
			i--
			continue
		} else if nums2[i] < nums1[k] {
			nums1[end] = nums1[k]
			k--
			end--
			continue
		} else {
			nums1[end] = nums1[k]
			end--
			nums1[end] = nums2[i]
			end--
			i--
			k--
			continue
		}
	}
	if i != -1 {
		for i >= 0 {
			nums1[i] = nums2[i]
			i--
		}
	}
	fmt.Println(nums1)
}
