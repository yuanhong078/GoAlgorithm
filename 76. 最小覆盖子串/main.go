package main

import (
	"fmt"
	"sort"
)

/*
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

说明：
	如果 S 中不存这样的子串，则返回空字符串 ""。
	如果 S 中存在这样的子串，我们保证它是唯一的答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(minWindow("a", "aa"))
}

func minWindow(s string, t string) string {


	if len(s)<len(t){
		return ""
	}

	idx := make([][]int, len(t))

	for k, v := range s {
		for k1, v2 := range t {
			if v == v2 {
				if idx[k1] == nil {
					idx[k1] = make([]int, 0)
					idx[k1] = append(idx[k1], k)
				} else {
					idx[k1] = append(idx[k1], k)
				}
			}
		}
	}

	for i := 0; i < len(t); i++ {
		if idx[i] == nil {
			return ""
		}
	}

	selected := make([]int, 0)

	sum := make(map[int][]int)
	for i := 0; i < len(idx[0]); i++ {
		minDistance(idx, 0, 0, idx[0][i], t, selected, sum)
	}
	min := len(s) * len(t)
	for k, v := range sum {
		fmt.Println("k:",k," v:",v)
		if k < min {
			min = k
		}
	}
	num := sum[min]
	sort.Ints(num)
	return s[num[0] : num[len(num)-1]+1]
}


func minDistance(idx [][]int, distance, i, point int, t string, selected []int, sum map[int][]int) {

	if i != len(t)-1 {
		for v := 0; v < len(idx[i+1]); v++ { //[B]int{}
			value := idx[i+1][v] //idx[B][0}
			//selected = append(selected, point) //idx[A][0} append
			d := getDistance(point, value)
			d2 := 0
			for _, v2 := range selected[:i] {
				d2 += getDistance(v2, value)
			}
			//fmt.Println("D:", distance+d+d2, " i:", i, " S:", append(selected, point), " V:", value)
			minDistance(idx, distance+d+d2, i+1, value, t, append(selected, point), sum)
		}
	} else {
		fmt.Println("sum:",distance," select:",append(selected, point))
		sum[distance] = append(selected, point)
	}

}

func getDistance(i, j int) int {
	if i < j {
		return j - i
	}
	return i - j
}
