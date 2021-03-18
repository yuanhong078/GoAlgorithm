package main

import "fmt"

/*
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	head := &ListNode{0, nil}
	head.Next = lists[0]

	for i := 1; i < len(lists); i++ {
		temp := lists[i]
		node := temp
		if head.Next == nil {
			head.Next = temp
			continue
		}
		for node != nil {
			if node.Val < head.Next.Val {
				c := node.Next
				node.Next = head.Next
				head.Next = node
				node = c
			} else {
				p := head.Next
				for p.Next != nil && p.Next.Val < node.Val {
					p = p.Next
				}

				//插入新的节点
				c := node.Next
				node.Next = p.Next
				p.Next = node
				node = c
			}
		}
	}
	return head.Next
}

func main() {
	lists := []*ListNode{
		//&ListNode{2, &ListNode{4, &ListNode{5, nil}}},
		//&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		//&ListNode{2, &ListNode{6, nil}},
		&ListNode{0, nil},
	}
	res := mergeKLists(lists)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}
