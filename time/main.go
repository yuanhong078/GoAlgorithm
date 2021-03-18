package main

import (
	"fmt"
	"time"
)

func main() {
	year := 2021
	month := time.Month(4)
	//使用尽可能短的代码获取任意月份的天数
	days := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	fmt.Println(days)
}
