package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"


示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。


说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println("res:", longestCommonPrefix([]string{"flower", "dlower"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := len(strs[0])
	longestWord := strs[0]
	for _, v := range strs {
		if len(v) < minLength {
			minLength = len(v)
			longestWord = v
		}
	}
	res := longestWord
	for i := 0; minLength > i; i++ {
		for _, v := range strs {
			if longestWord[i:1+i] == v[i:i+1] {
				continue
			} else {
				//if i==0{
				//	return ""
				//}
				res = longestWord[:i]
				return res
			}

		}
	}
	return res
}
