package main

import "fmt"

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	var head *ListNode
	node2 := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	node1 := &ListNode{Val: 2, Next: node2}
	head = &ListNode{Val: 1, Next: node1}
	res := reverseList(head)
	fmt.Println(res)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var next, cur, pre *ListNode
	pre = nil
	cur = head //1
	for cur.Next != nil {

		next = cur.Next //2
		cur.Next = pre  //1->nil
		pre = cur
		cur = next
		next = next.Next
		//if next.Next!=nil{
		//
		//}
		fmt.Println("p:",pre," n:",next," c:",cur)
	}
	cur.Next=pre
	head = cur
	//for head!=nil{
	//	head=head.Next
	//	fmt.Println("--n",head.Next)
	//}
	return head
}
