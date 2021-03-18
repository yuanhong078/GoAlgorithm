package main

import (
	"fmt"
	"math"
)

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	//res := divide(10, 3)
	//fmt.Println(res)
	//word := []string{"foo", "bar","zoo"}
	//temp := word
	//
	//ntemp := testArr(temp, "foo")
	//fmt.Println(ntemp, temp,word)


	temp := []string{"foo", "bar", "zoo"}

	for i := 0; i < len(temp); i++ {
		if temp[i] == "foo" {
			fmt.Println(temp[:i], "+++", temp[i+1:])

			ret := append(temp[:i], temp[i+1:]...)

			fmt.Println(temp[:i], "+++", temp[i+1:])

			fmt.Println("temp:", temp, "ret:",ret )
			break
		}
	}
	fmt.Println(temp)

	//arr:=[...]int{1,2,3,4,5,6,7}
	//s1:=arr[2:5]
	//s2:=s1[2:4]
	//fmt.Println("arr:",arr)
	//fmt.Println("s1:",s1)
	//fmt.Println("s2:",s2)
	//s2[0]=100
	//fmt.Println("arr:",arr)
	//fmt.Println("s1:",s1)
	//fmt.Println("s2:",s2)
	//s2=append(s2,s1...)
	//fmt.Println("arr:",arr)
	//fmt.Println("s1:",s1)
	//fmt.Println("s2:",s2)
	//s3:=append(s2,10)
	//s4:=append(s3,11)
	//s5:=append(s4,12)
	//fmt.Println("arr:",arr)
	//fmt.Println("s3,s4,s5:",s3,s4,s5)

}
func divide(dividend int, divisor int) int {
	res := 0
	if dividend == 0 {
		return res
	}

	isNegative := false
	if dividend > 0 && divisor < 0 {
		divisor = -divisor
		isNegative = true
	} else if dividend < 0 && divisor > 0 {
		dividend = -dividend
		isNegative = true
	} else if dividend < 0 && divisor < 0 {
		dividend, divisor = -dividend, -divisor
	}

	if divisor == 1 {
		res = dividend
	} else {
		for dividend >= divisor {
			dividend -= divisor
			res++
		}
	}

	if isNegative {
		res = -res
	}
	if res > int(math.Exp2(31)-1) || res < int(-math.Exp2(31)) {
		return int(math.Exp2(31) - 1)
	}
	return res

}

func testArr(temp []string, word string) []string {
	//for k, v := range temp {
	//	if word == v {
	//		ret := append(temp[:k], temp[k+1:]...)
	//		fmt.Println("bool:", true, " temp:", temp)
	//		return ret
	//	}
	//}

	for i := 0; i < len(temp); i++ {
		if temp[i] == word {
			r1 := temp[:i]
			r2 := temp[i+1:]
			fmt.Println(r1, r2)
			//ret := append(temp[0:i], temp[i+1:]...)
			fmt.Println("ptr:", "  ", &r2[0])
			ret := append(r1, r2...)
			fmt.Println(r1, r2, "=========")
			fmt.Println("ptr:", "  ", &r2[0])
			fmt.Println("bool:", true, " temp:", temp)
			return ret
		}
	}
	return nil
}
