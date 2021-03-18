package main

import (
	"fmt"
)

/*
给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。

示例 1:

输入: [[1,1],[2,2],[3,3]]
输出: 3
解释:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4


示例 2:

输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4
解释:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-points-on-a-line
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	points := [][]int{{0, 0}, {94911150, 94911151}, {94911151, 94911152}}
	fmt.Println(maxPoints(points))
}

type k struct {
	denominator int //分母
	numerator   int //分子
}

func maxPoints(points [][]int) int {
	if len(points) == 0 || len(points) == 2 {
		return len(points)
	}
	max := 0
	for i := 0; i < len(points); i++ {
		record := make(map[k]int)
		p := 0
		t := 0
		fmt.Println()
		for j := 0; j < len(points); j++ {
			if i != j {
				if points[i][0] == points[j][0] { //判断斜率是否存在
					if points[i][1] != points[j][1] { //判断两点是否重合
						p++
						if p+t > max {
							max = p + t
						}
						fmt.Println(i, j, max)
						continue
					} else {
						t++ //重合的点的个数
						fmt.Println(i, j, max)
						continue
					}
				}
				key:=gradient(points[i], points[j])
				record[key]++
				if record[key]+t > max {
					max = record[key] + t
				}
				fmt.Println(i, j, max, "--", key)
			}
		}
		if max == 0 { //只有重合情况下
			max += t
			fmt.Println(max, t, "==")
		}
	}
	return max + 1
}

//{0, 0}, {94911150, 94911151}, {94911151, 94911152}
func gradient(i, j []int) k {
	numerator := j[1] - i[1]
	denominator := j[0] - i[0]
	common := 0
	if numerator > denominator {
		common = gcd(numerator, denominator)
	} else {
		common = gcd(denominator, numerator)
	}
	res := k{denominator / common, numerator / common}
	return res
}

func gcd(i, j int) int {

	for j != 0 {
		i, j = j, i%j
	}

	return i
}
