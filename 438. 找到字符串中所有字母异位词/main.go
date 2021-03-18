package main

import (
	"fmt"
	"reflect"
)

/*
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：


	字母异位词指字母相同，但排列不同的字符串。
	不考虑答案输出的顺序。


示例 1:

输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/




func main() {
	res := findAnagrams("abacbabc", "abc")
	fmt.Println(res)
}

func findAnagrams1(s string, p string) []int {
	numbers := make(map[string]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(p); i++ {
		numbers[string(p[i])] += 1
	}

	//mNum := copyMap(numbers)
	mNum := make(map[string]int)

	for i := 0; i < len(s)-len(p)+1; i++ {

		isAna := true
		if i == 0 {
			t := i
			for t < i+len(p) {
				a := string(s[t])
				mNum[a]++
				t++
			}
		} else {
			mNum[string(s[i-1])]--
			mNum[string(s[i+len(p)-1])]++
		}

		for k, v := range numbers {
			if v2, ok := mNum[k]; !ok {
				isAna = false
				break
			} else if v != v2 {
				isAna = false
				break
			}
		}
		fmt.Println(i, mNum)
		if isAna {
			res = append(res, i)
		}
	}

	return res
}

func findAnagrams2(s string, p string) []int {
	numbers := make(map[string]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(p); i++ {
		numbers[string(p[i])] += 1
	}

	//mNum := copyMap(numbers)
	mNum := make(map[string]int)

	for i := 0; i < len(s)-len(p)+1; {
		isAna := true
		if i == 0 {
			t := i
			for t < i+len(p) {
				a := string(s[t])
				mNum[a]++
				t++
			}
		} else {
			a := string(s[i+len(p)-1])
			if _, ok := numbers[a]; !ok {
				i = i + len(p)
				if i+len(p) > len(s) {
					break
				}
				mNum = make(map[string]int)
				t := i
				for t < i+len(p) {
					b := string(s[t])
					mNum[b]++
					t++
				}
			} else {
				mNum[string(s[i-1])]--
				mNum[a]++
			}
		}
		for k, v := range numbers {
			if v2, ok := mNum[k]; !ok {
				isAna = false
				break
			} else if v != v2 {
				isAna = false
				break
			}
		}
		if isAna {
			res = append(res, i)
		}
		i++
	}
	return res
}

func findAnagrams(s string, p string) []int {
	numbers := make(map[string]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(p); i++ {
		numbers[string(p[i])] += 1
	}

	//mNum := copyMap(numbers)
	mNum := make(map[string]int)
	l, r := 0, 0

	for i := 0; i < len(s)-len(p)+1 && r < len(s); {
		a := string(s[r])
		//如果不是子串中的字母
		if _, ok := numbers[a]; !ok {
			i = r + 1
			l, r = i, i
			continue
		} else {
			if l == r {
				//为mNum赋值
				mNum = make(map[string]int)
				for r < l+len(p) {
					word := string(s[r])
					mNum[word]++
					r++
				}
				if reflect.DeepEqual(mNum, numbers) {
					res = append(res, l)
				}
				i++
			} else {
				mNum[string(s[l])]--
				mNum[a]++
				isAna := true
				for k, v := range numbers {
					if mNum[k] != v {
						isAna = false
						break
					}
				}
				if isAna {
					res = append(res, i)
				}
				i++
				l++
				r++
			}
		}

	}
	return res
}

func copyMap(source map[string]int) map[string]int {
	numbers := make(map[string]int, 0)
	for k, v := range source {
		numbers[k] = v
	}

	return numbers
}
