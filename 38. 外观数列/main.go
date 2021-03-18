package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
1.     1
2.     11
3.     21
4.     1211
5.     111221
*/

func main() {
	result := countAndSay(5)
	fmt.Println(result, "--")
}

func countAndSay(n int) string {

	var pre strings.Builder
	pre.WriteString("1")
	for i := 1; i < n; i++ {
		var s strings.Builder
		count := 0
		length:=pre.Len()
		word:=pre.String()
		for p := 0; p < length; p++ {
			count++
			if p+1 != length {
				if word[p] != word[p+1] {
					s.WriteString(strconv.Itoa(count) + string(word[p]))
					count = 0
				} else {
					continue
				}
			} else {
				s.WriteString(strconv.Itoa(count) + string(word[p]))
			}

		}

		pre = s
	}
	return pre.String()
}
