package main

import (
	"fmt"
	"reflect"
	"sync"
)

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true


示例 2:

输入: s = "rat", t = "car"
输出: false

说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main() {
	s:="anagram"
	t:="nagaram"

	fmt.Println(isAnagram(s,t))
}

func isAnagram(s string, t string) bool {
	if len(s)!= len(t){
		return false
	}
	m1:=make(map[uint8]int)
	for i:=0;i<len(s);i++{
		m1[s[i]]++
	}
	m2:=make(map[uint8]int)
	for i:=0;i<len(t);i++{
		m2[t[i]]++
	}

	return reflect.DeepEqual(m1, m2)
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	//创建两个仿bitmap的数组
	var wg sync.WaitGroup
	var wordstable_s =  [26]int{}
	var wordstable_t =  [26]int{}
	wg.Add(2)
	go func(){
		makewordtable(&wordstable_s,s)
		wg.Done()
	}()
	go func(){
		makewordtable(&wordstable_t,t)
		wg.Done()
	}()
	wg.Wait()
	//最后比较一下两个数组就好了
	return wordstable_s == wordstable_t

}

func makewordtable(wordtable *[26]int, s string) {
	for i:=0;i<len(s);i++{
		//使用asc2码 计算对应仿bitmap数组的下标
		index := s[i] - 'a'
		wordtable[index] ++
	}
}

