package main

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	length := 1
	p := head
	listArr := []*ListNode{}
	listArr = append(listArr, head)
	for p.Next != nil {
		p = p.Next
		listArr = append(listArr, p)
		length++
	}

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		if listArr[i].Val != listArr[j].Val {
			return false
		}
	}
	return true
}

func main() {

}
