package main

import "fmt"

/*
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	result := integerBreak(10)
	fmt.Println(result)
}

//8=2+3+3  18=2*3*3
func integerBreak(n int) int {
	record := make([]int, n+1)
	record[1] = 1
	record[2] = 1
	record[3] = 2
	record[4] = 4
	//record[5] = 6
	//record[5] = 6
	for i := 5; i <= n; i++ {

		record[i] =3* max(record[i-3],i-3)
		fmt.Println(i, " ", record[i])
	}
	return record[n]
}

func max(x, y int) int {
	if x>y {
		return x
	}
	return y
}
