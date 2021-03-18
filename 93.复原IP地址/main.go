package main

import "fmt"

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。


示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

示例 2：

输入：s = "0000"
输出：["0.0.0.0"]

示例 3：

输入：s = "1111"
输出：["1.1.1.1"]

示例 4：

输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]

示例 5：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	restoreIpAddresses( "0000")
}

func restoreIpAddresses(s string) []string {
	l := len(s)
	if l < 4 || l > 12 {
		return []string{}
	}

	//quotient := l / 4
	//remainder := l % 4

	res := splicIP(s, 0)
	for k, v := range res {
		fmt.Println("k: ", k, " ip: ", v)
	}
	return res

}

func splicIP(s string, i int) []string {
	if i == 3 {
		if len(s) == 0||len(s)>3 {
			return nil
		}
		if len(s) > 1 && s[0] == '0' {
			return nil
		}
		if len(s) == 3 && s > "255" {
			return nil
		}
		return []string{s}
	}
	k := 1
	ips := []string{}

	for (k < 4) {
		if len(s)-k+i < 3 {
			break
		}
		//开头为0
		if s[0] == '0' {
			res := splicIP(s[1:], i+1)
			if res != nil {
				for _, v := range res {
					ips = append(ips, s[0:k]+"."+v)
				}
			}
			break
		}
		//大于255
		if k == 3 && s[:k] > "255" {
			break
		}

		res := splicIP(s[k:], i+1)

		if res != nil {
			for _, v := range res {
				ips = append(ips, s[0:k]+"."+v)
			}
		}
		k++
	}

	if len(ips) == 0 {
		return nil
	}
	return ips
}
