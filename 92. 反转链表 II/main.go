package main

import "fmt"

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤m≤n≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
输出: 5->4->3->2->1->NULL

*/

func main() {
	var head *ListNode
	//node2 :=
	//node1 :=
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	res := reverseBetween(head, 1, 5)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	i := 1
	tail := head
	if m == 1 {
		next := head.Next
		for i != n {
			head, next.Next, next = next, head, next.Next
			i++
		}
		tail.Next = next
		return head
	}
	for tail != nil {
		if i+1 != m {
			tail = tail.Next
			i++
			continue
		}
		//i==m-1
		top := tail.Next //中间反转链表的头
		end := top
		next := top.Next

		for i != n-1 {
			top, next.Next, next = next, top, next.Next
			i++
		}
		tail.Next = top
		end.Next = next
		break
	}
	return head
}
