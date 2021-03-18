package main

import "fmt"

/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2,Next: nil}}}}}}
	res:=removeElements(head,4)
	for res!=nil{
		fmt.Println(res)
		res=res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	t := dummyHead

	for t.Next != nil {
		if t.Next.Val == val {
			t.Next = t.Next.Next
			continue
		}
		t = t.Next
	}
	return dummyHead.Next
}
