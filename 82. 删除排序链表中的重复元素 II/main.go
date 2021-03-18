package main

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

输入: 1->2->3->3->4->4->5
输出: 1->2->5

输入: 1->1->1->2->3
输出: 2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	t := dummyHead
	t.Next = head
	var p *ListNode
	for t.Next != nil {
		val := t.Next.Val
		p = t.Next.Next
		if p==nil{
			break
		}
		if p.Val != val {
			t = t.Next
			continue
		}
		for p != nil {
			if p.Val == val {
				p = p.Next
				continue
			}
			break
		}
		t.Next = p
	}
	return dummyHead.Next
}
