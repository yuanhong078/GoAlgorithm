package main

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：

你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//归并排序、快速排序
func sortList(head *ListNode) *ListNode {
	list := []int{}

	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	//排序算法

}
