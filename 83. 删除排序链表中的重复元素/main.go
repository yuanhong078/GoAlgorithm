package main

import "fmt"

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次
示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
 */

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	res := deleteDuplicates(head)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur:=head
	next:=cur.Next
	val:=cur.Val
	for next!=nil{
		fmt.Println(val,next.Val)
		if val!=next.Val{
			cur,next=next,next.Next
			val=cur.Val
			continue
		}
		//相等元素
		cur.Next,next=next.Next,next.Next
	}
	return head


}