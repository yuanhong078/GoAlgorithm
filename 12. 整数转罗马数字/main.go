package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
	I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。


给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println("res: ", intToRoman(1994))
}

func intToRoman2(num int) string {
	if num > 3999 || num < 1 {
		return ""
	}
	one, ten, hundred, thousand := getValue(num)
	res := ""
	for ; thousand > 0; thousand-- {
		res = res + "M"
	}

	for ; hundred > 0; hundred-- {
		if hundred == 9 {
			res = res + "CM"
			break
		}
		if hundred > 5 {
			res = res + "D"
			hundred -= 5
		}
		if hundred == 5 {
			res = res + "D"
			break
		}
		if hundred == 4 {
			res = res + "CD"
			break
		}
		res = res + "C"
	}
	for ; ten > 0; ten-- {
		if ten == 9 {
			res = res + "XC"
			break
		}
		if ten > 5 {
			res = res + "L"
			ten -= 5
		}
		if ten == 5 {
			res = res + "L"
			break
		}
		if ten == 4 {
			res += "XL"
			break
		}
		res += "X"
	}
	for ; one > 0; one-- {
		if one == 9 {
			res += "IX"
			break
		}
		if one == 5 {
			res += "V"
			break
		}
		if one == 4 {
			res += "IV"
			break
		}
		if one > 5 {
			res += "V"
			one -= 5
		}
		res += "I"
	}
	return res
}

func getValue(num int) (one int, ten int, hundred int, thousand int) {
	one, num = num%10, num/10
	ten, num = num%10, num/10
	hundred, num = num%10, num/10
	thousand = num
	return
}

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := ""
	x := len(intSlice) - 1
	for ; num != 0; {
		i := len(strconv.Itoa(num)) - 1
		n := (num / int(math.Pow10(i))) * int(math.Pow10(i))	//最高位的数值
		for ; x >= 0; x-- {			//200
			if n >= intSlice[x] {
				s += roman[x]
				num -= intSlice[x]
				break
			}
		}
	}
	return s
}

