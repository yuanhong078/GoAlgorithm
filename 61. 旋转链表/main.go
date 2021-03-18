package main

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	front := &ListNode{}
	front.Next = head
	dummyHead, end := front, front
	k2, length := 0, 0

	//遍历数组，使front指向k前节点，end指向末尾节点
	for i := 0; i < k; i++ {
		if end.Next == nil {
			k2 = k - i //i==3,end==2,k2=1
			length = i

			break
		}
		end = end.Next
	}

	//链表的长度超过K
	if k2 != 0 {
		p := k2 % length
		if p == 0 {
			return head
		}
		end = front
		for i := 0; i < p; i++ {
			end = end.Next
		}

		for end.Next != nil {
			front = front.Next
			end = end.Next
		}
		dummyHead = front.Next
		front.Next = nil
		if end != head {
			end.Next = head
		} else {
			end.Next = nil
		}
	}

	//链表的长度不超过K
	if k2 == 0 {
		for end.Next != nil {
			front, end = front.Next, end.Next
		}

		dummyHead = front.Next
		front.Next = nil
		if end != head && dummyHead != head {
			end.Next = head
		} else {
			end.Next = nil
		}

	}
	return dummyHead
}

func main() {

}
