package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321

 示例 2:

输入: -123
输出: -321

示例 3:

输入: 120
输出: 21


*/

func main() {
	fmt.Println(reverse(1563847412))
}

func reverse(x int) int {
	if x == 0 {
		return x
	}
	strx := strconv.Itoa(x)
	arrx := []rune(strx)
	i := 0

	if x < 0 {
		i = 1 //符号位
	}

	for j := len(strx) - 1; i < j; i, j = i+1, j-1 {

		arrx[i], arrx[j] = arrx[j], arrx[i]
	}

	for {
		if arrx[0] == '0' {
			arrx = arrx[1:]
		} else {
			break
		}
	}

	result, err := strconv.Atoi(string(arrx))

	if err != nil {
		fmt.Println("array conver int err :", err)
	}
	fmt.Println(x)
	fmt.Println(int(math.Exp2(31))-1)
	if result>(int(math.Exp2(31))-1) || result<int(-math.Exp2(31)) {
		return 0
	}

	return result
}
