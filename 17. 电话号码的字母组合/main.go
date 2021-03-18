package main

import (
	"fmt"
)

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

func main() {
	fmt.Println(letterCombinations1("23"))

}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	nums := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	arr := []rune(digits)
	//order := make(map[int][]string)
	//for k, v := range arr {
	//	order[k] = nums[string(v)]
	//}
	length := len(arr)
	if length == 1 {
		return nums[string(arr[0])]
	}
	temp := make([]string, 0)
	temp = append(temp, nums[string(arr[0])]...)

	for i := 1; i < length; i++ {
		tL := len(temp)
		for j := 0; j < tL; j++ {
			num := string(arr[i])
			for _, v := range nums[num] {
				temp = append(temp, temp[j]+v)
			}

		}
		temp = temp[tL:]
	}
	return temp
}

func letterCombinations1(digits string) []string {
	if digits == "" {
		return nil
	}

	var start string

	res := findCombination(digits, 0, start)
	return res
}

func findCombination(digits string, index int, s string) (res []string) {
	Letter := map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	if index == len(digits) {
		fmt.Println("s:", s)
		res = append(res, s)
		return
	}
	c := digits[index]

	letters := Letter[string(c)]
	for i := 0; i < len(letters); i++ {
		fmt.Println("num:", string(letters[i]))
		r:=findCombination(digits, index+1, s+string(letters[i]))
		res =append(res,r...)
	}
	return
}
