package main

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。
进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//7 2 4 3
//  5 6 4
//7 8 0 7

//	应考虑数字过大，越界的问题
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	}
	var num1, num2 []int

	for l1 != nil || l2 != nil {
		if l1 != nil {
			num1 = append(num1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = append(num2, l2.Val)
			l2 = l2.Next
		}
	}

	head := &ListNode{}

	if len(num2) > len(num1) {
		num2, num1 = num1, num2 //保存num1的长度是最大的
	}

	next, val := 0, 0

	for i, j := len(num1)-1, len(num2)-1; i >= 0 ; { //头插法
		if j != -1 {
			val = num1[i] + num2[j] + next
			i--
			j--
		} else {
			val = num1[i] + next
			i--
		}
		next = val / 10
		val1:=val%10
		node := &ListNode{Val: val1}
		head.Next, node.Next = node, head.Next
	}
	if next!=0{
		node := &ListNode{Val: 1}
		node.Next=head.Next
		return node
	}
	return head.Next
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {	//两个数子长度相等
//	//var head *ListNode	为nil
//	head := &ListNode{}	   //不为nil
//	t := head
//	for l1 != nil || l2 != nil {
//
//		if t == nil { //考虑高位进1
//			v1, v2 := 0, 0
//			n1, n2 := 0, 0
//			if l1 != nil {
//				v1 = l1.Val
//				if l1.Next != nil {
//					n1 = l1.Next.Val
//				}
//				l1 = l1.Next
//			}
//			if l2 != nil {
//				v2 = l2.Val
//				if l2.Next != nil {
//					n2 = l2.Next.Val
//				}
//				l2 = l2.Next
//			}
//			r := (n1 + n2) / 10
//			num := v1 + v2 + r
//			t.Val = num / 10                          //1-->
//			t.Next = &ListNode{num % 10, &ListNode{}} //1-->0-->
//			t = t.Next
//		} else { //非头节点，只考虑低位
//			v1, v2 := 0, 0
//			n1, n2 := 0, 0
//			if l1 != nil {
//				v1 = l1.Val
//				if l1.Next != nil {
//					n1 = l1.Next.Val
//				}
//				l1 = l1.Next
//			}
//			if l2 != nil {
//				v2 = l2.Val
//				if l2.Next != nil {
//					n2 = l2.Next.Val
//				}
//				l2 = l2.Next
//			}
//			r := (n1 + n2) / 10
//			num := v1 + v2 + r
//			t.Next = &ListNode{num % 10, nil}
//			t = t.Next
//		}
//	}
//	return head
//}
