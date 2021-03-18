package main

import "fmt"

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

*/

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2,Next: nil}}}}}}
	res := partition(head, 3)
	myPrint(res)
}
func myPrint(res *ListNode){
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	smallHead, bigHead := &ListNode{0, nil}, &ListNode{0, nil}
	s, b := smallHead, bigHead
	for head != nil {
		if head.Val < x {
			s.Next = head
			s = s.Next
			head = head.Next
			continue
		}
		b.Next = head
		b = b.Next
		head = head.Next
		continue
	}
	b.Next=nil
	s.Next = bigHead.Next
	return smallHead.Next

}
