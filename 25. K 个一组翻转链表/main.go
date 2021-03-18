package main

import "fmt"

/*
难度困难614收藏分享切换为英文关注反馈给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	node := head
	res := &ListNode{}
	p := res
	for node != nil {
		isK := true
		list := []*ListNode{}
		for i := 0; i < k; i++ {
			list = append(list, node)
			if node.Next == nil && i+1 != k {
				isK = false
				node = nil
				break
			}
			node = node.Next
		}

		//对可以整除的部分进行反转
		if isK {
			for i := k - 1; i >= 0; i-- {
				if i == 0 && node == nil { //结尾节点的next值为nil
					list[i].Next = nil
				}
				p.Next = list[i]
				p = p.Next
			}
		} else {
			for i := 0; i < len(list); i++ {
				p.Next = list[i]
				p = p.Next
			}
		}
	}

	return res.Next
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	res := reverseKGroup(list, 5)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}
