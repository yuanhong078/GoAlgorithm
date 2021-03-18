package main

import (
	"fmt"
	"math"
)

/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	res := numSquares(12)
	fmt.Println(res)
}

func numSquares(n int) int {
	mem := make([]int, n+1)

	for i := 1; i <= n; i++ {

		a := math.Sqrt(float64(i))
		acc := int(a) * int(a)
		//fmt.Println("acc:",acc,i,a)
		if acc == i {
			mem[i] = 1
			//fmt.Println(mem[i],"  -")
		} else {
			t := i
			for j:=1;j<i/2+1;j++{
				 t=min(mem[j]+mem[i-j],t)
			}
			mem[i] = t
			//fmt.Println(mem[i],"  =")
		}
		fmt.Println(i, "  ", mem[i])
	}

	return mem[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
