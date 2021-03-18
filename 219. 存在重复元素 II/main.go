package main

import "fmt"

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{1,2,3,1,2,3}
	k:=2
	fmt.Println(containsNearbyDuplicate(nums,k))
}
//数组里保存的元素的个数不能超过k,i插入后，就要删除头元素，维护map的长度
func containsNearbyDuplicate(nums []int, k int) bool {

	duplicateMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _,ok:=duplicateMap[nums[i]];ok{
			return true
		}
		duplicateMap[nums[i]] = true
		//保持map里最多k个元素
		if len(duplicateMap) == k+1 {
			delete(duplicateMap, nums[i-k])
		}
	}
	return false
}
