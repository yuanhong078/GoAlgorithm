package main

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrderBottom(root *TreeNode) [][]int {

	res=make([][]int,0)

	if root==nil{
		return nil
	}
	a:=dfs(root,1)
	if len(a)>0{
		res=append(res,a)
	}

	res=append(res,[]int{root.Val})

	return res
}

func dfs(root *TreeNode,  level int) []int {
	if root == nil {
		return nil
	}
	res = append(res, []int{})
	left := dfs(root.Left, level+1)
	right := dfs(root.Right,  level+1)
	merge := []int{}
	merge = append(merge, left...)
	merge = append(merge, right...)

	if len(merge)>0{
		res[len(res)-level-1] = append(res[len(res)-level-1], merge...)
	}


	a := []int{}
	if root.Left != nil {
		a = append(a, root.Left.Val)
	}
	if root.Right != nil {
		a = append(a, root.Right.Val)
	}

	return a
}
