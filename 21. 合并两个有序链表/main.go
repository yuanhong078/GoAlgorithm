package main

import "fmt"

/*
难度简单1125收藏分享切换为英文关注反馈将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	head := newList
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {

		if l1.Val >= l2.Val {
			newList.Val = l2.Val
			fmt.Println("v2: ", newList.Val)
			newList.Next = &ListNode{}
			newList = newList.Next
			l2 = l2.Next
		} else {
			newList.Val = l1.Val
			//fmt.Println("v1: ", newList.Val)
			newList.Next = &ListNode{}
			newList = newList.Next
			l1 = l1.Next
		}
	}

	if l1!=nil{
		newList.Val = l1.Val
		newList.Next = l1.Next
	}else if l2 != nil {
		newList.Val = l2.Val
		newList.Next = l2.Next
	}

	return head
}

func main() {
	var l1, l2 *ListNode

	//l1 = &ListNode{Val: 1, Next: nil}
	//l2 = &ListNode{Val: 2, Next: nil}
	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{4, nil}}}
	l2 = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{4, nil}}}
	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Println("val: ", res.Val)
		res = res.Next
	}

}
