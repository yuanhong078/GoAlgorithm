package main

import (
	"fmt"
	"strings"
)

/*
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true

示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false

示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false

示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, s string) bool {
	wordMap := make(map[uint8]string)

	ss := strings.Split(s, " ")

	if len(ss) != len(pattern) {
		return false
	}

	for k, word := range ss {
		if v, ok := wordMap[pattern[k]]; ok {
			if v != word {
				return false
			}
		}
		wordMap[pattern[k]] = word
	}

	for k, v := range wordMap {
		for k2, v2 := range wordMap {
			if k != k2 && v == v2 {
				return false
			}
		}
	}

	return true
}
