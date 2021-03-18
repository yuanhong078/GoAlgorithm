package main

import (
	"fmt"
)

/*
二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。

每个 LED 代表一个 0 或 1，最低位在右侧。

例如，上面的二进制手表读取 “3:25”。

给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。

示例：

输入: n = 1
返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

提示：

	输出的顺序没有要求。
	小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
	分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
	超过表示范围（小时 0-11，分钟 0-59）的数据将会被舍弃，也就是说不会出现 "13:00", "0:61" 等时间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-watch
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	r := readBinaryWatch(2)
	fmt.Println(";;",r)
}

//func readBinaryWatch(num int) []string {
//	hour := []int{1, 2, 4, 8}
//	minute := []int{1, 2, 4, 8, 16, 32}
//	p := make([]string, 0)
//	p = generateWatch(num, hour, minute)
//	return p
//}
//
//func generateWatch(num int, hour, minute []int) []string {
//	p := make([]string, 0)
//
//	for i := 0; i <= num; i++ {
//		if i > 3 {
//			break
//		}
//
//		resHour := make([]int, 0)
//		resMin := make([]int, 0)
//
//		resHour = generateHour(0, i, 0, hour, resHour)
//
//		resMin = generateMin(0, num-i, 0, minute, resMin)
//
//		for _, h := range resHour {
//			var s strings.Builder //字符串拼接
//			sh := strconv.Itoa(h)
//
//			for _, m := range resMin {
//				s.WriteString(sh)
//				if m < 10 {
//					sm := strconv.Itoa(m)
//					s.WriteString(":0")
//					s.WriteString(sm)
//				} else {
//					sm := strconv.Itoa(m)
//					s.WriteString(":")
//					s.WriteString(sm)
//				}
//				p = append(p, s.String())
//				s.Reset()
//			}
//		}
//	}
//
//	return p
//}
//
//func generateHour(start, n, sum int, hour, res []int) []int {
//
//	if n == 0 {
//		res = append(res, sum)
//		return res
//	}
//
//	for i := start; i < 4; i++ {
//		sum += hour[i]
//		if sum > 11 {
//			return res
//		}
//		res = generateHour(i+1, n-1, sum, hour, res)
//		sum -= hour[i]
//	}
//
//	return res
//}
//
//func generateMin(start, n, sum int, minute, res []int) []int {
//	if n == 0 {
//		res = append(res, sum)
//
//		return res
//	}
//
//	for i := start; i < 6; i++ {
//		sum += minute[i]
//		if sum > 59 {
//			return res
//		}
//		res = generateMin(i+1, n-1, sum, minute, res)
//		sum -= minute[i]
//	}
//
//	return res
//}

//=================================================

//func readBinaryWatch(num int) []string {
//	hours := [4]int{1,2,4,8}
//	minutes := [6]int{1,2,4,8,16,32}
//
//	res := make([]string,0)
//	//接受转换后的字符串
//	s := ""
//	//计算小时
//	shour := 0
//	//计算分钟
//	sminute := 0
//
//	if num == 0{
//		s = "0:00"
//		res = append(res,s)
//		return res
//	}
//
//	Run(hours,minutes,s,&res,shour,sminute,0,0,num)
//
//	return res
//}
//
//func Run(hours [4]int,minutes [6]int,s string,res *[]string,shour,sminute,h,m,num int) {
//	if num == 0 {
//		s += strconv.Itoa(shour)
//		s += ":"
//		if sminute <=9 {
//			s += "0"
//		}
//		s += strconv.Itoa(sminute)
//		*res = append(*res,s)
//		return
//	}
//
//	//分钟
//	for i:=m;i<6;i++ {
//		if  sminute+minutes[i] <= 59 {
//			//再回溯回来的时候m之前不再进行选择
//			//因为已经选择并判断过
//			m++
//			Run(hours,minutes,s,res,shour,sminute+minutes[i],h,m,num-1)
//		}
//		fmt.Println("m1:",m,sminute)
//	}
//	fmt.Println(sminute)
//	fmt.Println("m:",m)
//	//小时
//	for i:=h;i<4;i++ {
//		if  shour+hours[i] <= 11 {
//			h++
//			Run(hours,minutes,s,res,shour+hours[i],sminute,h,m,num-1)
//		}
//	}
//	fmt.Println("H:",shour)
//}

func readBinaryWatch(num int) []string {
	res := []string{}
	// 假想 10个 bit 灯 index  --->       h1=>0   h2=>1   h4=>2   h8=>3   m1=>4   m2=>5   m4=>6   m8=>7   m16=>8   m32=>9
	//start = 0 回溯从bit 的 index = 0 开始, 10 代表时分灯的总数, 相当于10个灯中选中num个亮的灯
	backTrack(0, 10, num, []int{}, &res)
	return res
}

//backTrack 回溯算法 cap 灯的总数量, target 目标亮灯的数量 path []int  亮灯的index集合(记录回溯的路径) res 指针返回结果
func backTrack(start, cap, target int, path []int, res *[]string) {
	if len(path) == target { //选择亮灯的数量达到target
		//min, hour := 0, 0
		var min, hour uint
		for _, v := range path {

			if v >= 4 { // 如果灯的下表超过3 则表示这些灯是 minute的灯
				min += 1 << (v - 4)
				fmt.Println(v)
			} else { //如果 灯的index 小于4 表示是 hour的灯亮
				hour += 1 << v
			}
		}
		if min < 60 && hour < 12 { //排除不符合规则的,判断min hour 值是否有效
			*res = append(*res, fmt.Sprintf("%d:%02d", hour, min)) //格式化拼接成字符串
		}
		path = []int{} //重置 回溯的path
		return
	}
	for i := start; i < cap; i++ {
		path = append(path, i)
		backTrack(i+1, cap, target, path, res) //注意这里是 i+1
		path = path[:len(path)-1]              //对 path 进行rollback
	}
}
