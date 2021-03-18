package main

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	level int
	*TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []Node{}
	node := Node{0, root}
	queue = append(queue, node)
	s := []int{root.Val}
	res = append(res, s)

	for len(queue) > 0 { //层序遍历	que==nil []
		v := queue[0]
		level := v.level + 1 //新一层
		e := v.TreeNode

		if e.Left != nil {
			queue = append(queue, Node{level, e.Left}) //入队
			if len(res) < level+1 {
				s := make([]int, 0)
				res = append(res, s)
			}
			res[level] = append(res[level], e.Left.Val)
		}

		if e.Right != nil {
			queue = append(queue, Node{level, e.Right})
			if len(res) < level+1 {
				s := make([]int, 0)
				res = append(res, s)
			}

			res[level] = append(res[level], e.Right.Val)
		}
		queue = queue[1:]
	}

	return res
}
