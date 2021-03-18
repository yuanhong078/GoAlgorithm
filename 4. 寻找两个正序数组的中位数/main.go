package main

import "fmt"

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0

*/
func main() {
	result := findMedianSortedArrays([]int{0, 1}, []int{2, 3})
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) == 0 {
		return getMedian(nums2)
	}
	if len(nums2) == 0 {
		return getMedian(nums1)
	}
	sortedArr := getSortedArray(nums1, nums2)

	return getMedian(sortedArr)
}

func getMedian(nums1 []int) float64 {
	length := len(nums1)
	if length%2 == 0 {
		index2 := length / 2
		index1 := index2 - 1
		result := float64(nums1[index1]+nums1[index2]) / float64(2)
		return result
	}

	return float64(nums1[length/2])
}

func getSortedArray(nums1 []int, nums2 []int) []int {
	sortedArray := make([]int, 0)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if !(nums1[i] > nums2[j]) {
			sortedArray = append(sortedArray, nums1[i])
			i++
		} else {
			sortedArray = append(sortedArray, nums2[j])
			j++
		}
		if i == len(nums1) {
			sortedArray = append(sortedArray, nums2[j:]...)
		} else if j == len(nums2) {
			sortedArray = append(sortedArray, nums1[i:]...)
		}
	}

	return sortedArray

}
