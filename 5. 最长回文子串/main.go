package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：

输入: "cbbd"
输出: "bb"

*/
func main() {
	s := "a"
	s1 := longestPalindrome(s)
	fmt.Println("result: ",s1)
}

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}
	longestP := ""
	for i := 0; i < len(s)-1;i++ {
		j := i
		for j < len(s) {
			var isPalindrome bool
			if s[j] == s[i] { //判断是否为回文
				start := i
				end := j
				isPalindrome=true
				for start < end {
					if s[start] != s[end] {
						isPalindrome = false
						break
					}
					start += 1
					end -= 1
				}
			}
			if isPalindrome {
				if j+1-i> len(longestP){
					longestP = s[i:j+1]
					fmt.Println("new string: ",longestP)
				}
			}
			j++
		}
	}
	return longestP
}
