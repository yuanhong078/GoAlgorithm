package main

/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	res = preOrder(root, res)
	return res
}

func preOrder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = preOrder(root.Left, res)
	res = preOrder(root.Right, res)
	return res

}

func main() {

}
