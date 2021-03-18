package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

输入：(5) + (5 )
输出：0->1
*/
func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	res := addTwoNumbers(l1, l2)
	for res != nil {
		fmt.Println(*res, " =")
		res = res.Next
	}
	fmt.Println("===================================")
	l1 = &ListNode{9, &ListNode{2, nil}}
	l2 = &ListNode{3, &ListNode{7, nil}}
	res = addTwoNumbers(l1, l2)
	for res != nil {
		fmt.Println(*res, " =")
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	//res.add(l1, l2)
	temp := res
	index := 0
	for l1 != nil || l2 != nil {
		node := &ListNode{}
		if l1 != nil && l2 != nil {
			node.Val = (l1.Val + l2.Val + index) % 10
			index = (l1.Val + l2.Val + index) / 10

		} else if l1 == nil && l2 != nil {
			node.Val = (l2.Val + index) % 10
			index = (index + l2.Val) / 10
		} else if l1 != nil && l2 == nil {
			node.Val = (l1.Val + index) % 10
			index = (index + l1.Val) / 10
		}

		temp.Next = node //如果写成temp=node 则错误，第一个temp 应该指向root
		temp = temp.Next

		//l1, l2 = l1.Next, l2.Next

		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	if index != 0 {
		temp.Next = &ListNode{index, nil}
	}
	return res.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	headList := new(ListNode)
	head := headList
	num := 0

	for (l1 != nil || l2 != nil || num > 0) {
		headList.Next =  new(ListNode)
		headList = headList.Next
		if l1 != nil {
			num = num + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num = num + l2.Val
			l2 = l2.Next
		}
		headList.Val = (num) % 10
		num = num / 10
	}

	return head.Next
}
