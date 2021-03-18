package main

/*
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
     X  0  X  0  X
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	res := oddEvenList(head)
	for res != nil {
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	odd, even := &ListNode{}, &ListNode{}
	o, e := odd,even
	i := 0
	for head != nil {
		if (i+1)%2 == 1 { //奇数节点
			o.Next = head
			o, head = o.Next, head.Next
			i++
		} else { //偶数节点
			e.Next = head
			e, head = e.Next, head.Next
			i++
		}
	}
	e.Next = nil
	o.Next = even.Next
	return odd.Next
}
