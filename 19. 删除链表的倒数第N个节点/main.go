package main

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
					 n
给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.


说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {

	length := 1
	list := head
	for list.Next != nil {
		length++
		list = list.Next
	}

	k := length - n
	if k == 0 {
		return head.Next
	}
	node := head
	for i := 0; i < k; i++ {
		if i+1 == k {
			temp := node.Next
			node.Next = temp.Next
			break
		}
		node = node.Next
	}

	return head
}

//双指针发求解
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	front, end := &ListNode{}, &ListNode{}
	front.Next, end.Next = head, head
	t:=front

	for i := 0; i < n; i++ {
		end = end.Next
	}

	for end.Next != nil {
		front = front.Next
		end = end.Next
	}

	//front指向被删除的前一个节点，end 指向最后的节点
	if front.Next!=nil{
		front.Next=front.Next.Next
	}
	return t.Next

}

func main() {

}
