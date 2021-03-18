package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("LVIII"))
}

func romanToInt(s string) int {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intValue := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	res := 0
	length := len(s)
	for length > 0 {
		for i := len(roman) - 1; i >= 0; i-- {
			if strings.HasPrefix(s, roman[i]) {
				res += intValue[i]
				s = strings.TrimPrefix(s, roman[i])
				length = len(s)
				break
			}
		}
	}
	return res
}
