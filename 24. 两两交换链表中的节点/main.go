package main

import "fmt"

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	temp := head
	for temp.Next != nil {
		temp.Val, temp.Next.Val = temp.Next.Val, temp.Val
		if temp.Next.Next != nil {
			temp = temp.Next.Next
		} else {
			break
		}
	}
	return head
}

func main() {

	head := &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{2, nil}}}}
	res := swapPairs(head)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}
