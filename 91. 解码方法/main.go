package main

import "fmt"

/*
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

题目数据保证答案肯定是一个 32 位的整数。

 

示例 1：

输入：s = "12"
输出：2
解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2：

输入：s = "226"
输出：3
解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
示例 3：

输入：s = "0"
输出：0
示例 4：

输入：s = "1"
输出：1
示例 5：

输入：s = "2"
输出：1
 

提示：

1 <= s.length <= 100
s 只包含数字，并且可能包含前导零。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-ways
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	res := numDecodings("1111112")
	fmt.Println(res)
}

func numDecodings(s string) int {

	memo := make(map[string]int)
	res := numRecord(s, memo)
	return res
}

func numRecord(num string, memo map[string]int) int {
	if v, ok := memo[num]; ok {
		fmt.Println("=====")
		return v
	}
	//递归终止条件
	if len(num) == 1 {
		if num == "0" {
			return 0
		}
		memo[num] = 1
		return 1
	}
	if len(num) == 0 {
		return 0
	}

	a1 := []byte(num[:2])
	t1, t2 := 0, 0

	if a1[0] == '0' {
		memo[num] = t1 + t2
		return 0
	} else if a1[0] < '3' {
		if a1[0] == '2' {
			if a1[1] <= '6' {
				if len(num) > 2 {
					t1 = numRecord(num[2:], memo)
				} else {
					t1 = 1
				}
			}
		} else {
			if len(num) > 2 {
				t1 = numRecord(num[2:], memo)
			} else {
				t1 = 1
			}
		}
	}

	t2 = numRecord(num[1:], memo)
	memo[num] = t1 + t2

	return t1 + t2
}
