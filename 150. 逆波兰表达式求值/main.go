package main

import (
	"fmt"
	"strconv"
)

/*
根据 逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

 

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
 

示例 1：

输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/evaluate-reverse-polish-notation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			length := len(stack)
			if length < 2 {
				return 0
			}
			res := stack[length-2] + stack[length-1]
			stack = append(stack[:length-2], res)
		} else if tokens[i] == "-" {
			length := len(stack)
			if length < 2 {
				return 0
			}
			res := stack[length-2] - stack[length-1]
			stack = append(stack[:length-2], res)
		} else if tokens[i] == "*" {
			length := len(stack)
			if length < 2 {
				return 0
			}
			res := stack[length-2] * stack[length-1]
			stack = append(stack[:length-2], res)
		} else if tokens[i] == "/" {
			length := len(stack)
			if length < 2 {
				return 0
			}
			res := stack[length-2] / stack[length-1]
			stack = append(stack[:length-2], res)
		} else {
			num,err:=strconv.Atoi(tokens[i])
			if err!=nil{
				return -1
			}
			stack = append(stack, num)
		}

	}

	return stack[0]
}

func main() {
	stack := []int{1, 2, 3}
	res := 5
	stack = append(stack[:1], res)
	fmt.Println(stack)
}
