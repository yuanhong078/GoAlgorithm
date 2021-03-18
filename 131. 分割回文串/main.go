package main

import (
	"fmt"
	"reflect"
)

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	partition("ababbbabbaba")
}

func partition(s string) [][]string {
	res := split(s)

	result := [][]string{res[0]}

	for _, v := range res {
		repeat := false
		for _, v2 := range result {
			if reflect.DeepEqual(v, v2) {
				repeat = true
				break
			}
		}
		if !repeat {
			result = append(result, v)
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}
	return result
}

func split(s string) [][]string {
	if len(s) == 1 {
		return [][]string{[]string{s}}
	}

	res := [][]string{}

	if isPalindrome(s) {
		p := []string{s}
		res = append(res, p)
	}
	i := 1
	for (i < len(s)) {
		r1 := split(s[:i])
		r2 := split(s[i:])
		if r1 == nil || r2 == nil {
			continue
		}

		for _, v1 := range r1 {
			for _, v2 := range r2 {
				e := []string{}
				e = append(e, v1...)
				e = append(e, v2...)
				res = append(res, e)
			}
		}
		i++
	}
	if len(res) != 0 {
		return res
	}
	return nil
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
