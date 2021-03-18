package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true


示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	result := isPalindrome("0P")
	fmt.Println(result)

}

func isPalindrome1(s string) bool {
	if s == "" {
		return true
	}
	i := 0
	j := len(s) - 1

	for i < j {
		a := strings.ToLower(string(s[i]))
		if a < "0" || a > "z" || (a > "9" && a < "a") {
			i++
			continue
		}

		b := strings.ToLower(string(s[j]))
		if b < "0" || b > "z" || (b > "9" && b < "a") {
			j--
			continue
		}
		if a != b {
			return false
		} else {
			i++
			j--
		}
	}
	return true

}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	i := 0
	j := len(s) - 1

	for i < j {
		a := strings.ToLower(string(s[i]))
		if !(unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i]))) {
			i++
			continue
		}

		b := strings.ToLower(string(s[j]))
		if !(unicode.IsLetter(rune(s[j])) || unicode.IsDigit(rune(s[j]))) {
			j--
			continue
		}
		if a != b {
			return false
		} else {
			i++
			j--
		}
	}
	return true

}
