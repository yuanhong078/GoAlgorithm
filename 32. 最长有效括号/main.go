package main

import "fmt"

/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"


示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"


*/

func main() {
	res := longestValidParentheses("()(()")
	fmt.Println("res:", res)
}
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	var arr []string
	res := 0
	for i := 0; i < len(s); i++ {
		arr = append(arr, string(s[i]))
		length := len(arr)

		if length < 2 {
			continue
		} else {
			if arr[length-1] == ")" {
				if arr[length-2] == "(" {
					res += 2
					arr = arr[:length-2]
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
	res1 := len(s) - len(arr)
	fmt.Println("res1:", res1,len(arr))

	return res
}
