package main

import (
	"fmt"
	"strings"
)

/*
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func main() {
	maxLength := lengthOfLongestSubstring2("abcaa")
	fmt.Println(maxLength)

}

func lengthOfLongestSubstring(s string) int {
	max := 0
	index := make(map[rune]int)
	last := 0
	for i, e := range []rune(s) {
		if first, ok := index[e]; ok != false {
			if !(last > first) {
				last = first + 1 //新的不重复子串的开头
			}
		}
		if i-last+1 > max {
			max = i - last + 1
		}
		index[e] = i
	}

	return max
}
func lengthOfLongestSubstring2(s string) int {
	max := 0

	last := 0
	for i, _ := range s {
		index := strings.IndexByte(s[last:i], s[i])
		if !(index < last) {
			last = index + 1
		}
		if i-last+1 > max {
			max = i - last + 1
		}

	}

	return max
}
