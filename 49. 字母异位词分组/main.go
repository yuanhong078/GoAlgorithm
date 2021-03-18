package main

import "fmt"

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

说明：


	所有输入均为小写字母。
	不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	words := make(map[[26]int][]int)

	for k, v := range strs {
		s := []byte(v)
		str := [26]int{}
		for _, e := range s {
			str[e-'a'] += 1
		}
		words[str]=append(words[str],k)
	}

	res:=[][]string{}
	for _,v:=range words{
		s:=[]string{}
		for _,t:=range v{
			s=append(s,strs[t])
		}
		res = append(res,s)
	}
	return res
}
