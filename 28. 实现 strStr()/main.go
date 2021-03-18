package main

import "fmt"

/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	res := strStr("hello", "bll")
	fmt.Println(res)
}
func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}
	lN := len(needle)
	lH := len(haystack)
	for i := 0; i < lH; i++ {
		if i+lN > lH{
			return -1
		}
		if haystack[i] == needle[0] {
			t := i
			j := 0
			isSub := true
			for t < lH && j < lN {
				if haystack[t] != needle[j] {
					isSub = false
					break
				}
				t++
				j++
			}

			if j == lN && isSub {
				return i
			}
		}
	}
	return -1
}
