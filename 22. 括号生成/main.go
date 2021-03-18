package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	res := generateParenthesis(5)
	//for k, v := range res {
	//	fmt.Println("k:", k, " v:", v)
	//}
	res1 := []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"}



	for k1, v := range res1 {
		exist := false
		for k2, v2 := range res {
			if v == v2 {
				exist = true
				fmt.Println("k1:", k1, " v:", v, " k2:", k2, " v2:", v2)
				break
			}
		}
		if !exist {
			fmt.Println("not exist: ", v)
		}
	}
	fmt.Println("L1:", len(res1), " L2:", len(res))
}

var allP map[int][]string
func generateParenthesis(n int) []string {
	// 使用new方法开辟内存返回内存地址
	res := new([]string)

	backtracking(n, n, "",  res)

	return *res
}

func backtracking(left, right int, tmp string, res *[]string) {
	/*
	   回溯跳出条件，
	   并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
	   所以右括号用完了就意味着左括号必定用完了
	*/
	if right == 0 {
		*res = append(*res, tmp)
		return
	}

	// 生成左括号
	if left > 0 {
		backtracking(left - 1, right, tmp + "(", res)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		backtracking(left, right - 1, tmp + ")", res)
	}
}


//func generateParenthesis(n int) []string {
//
//	allP = make(map[int][]string, 0)
//
//	res := make([]string, 0)
//	res = g(n)
//	return res
//
//}
//
//func g(n int) []string {
//	res := make([]string, 0)
//	res = append(res, generateP(n)...)
//	for i := 2; i < n; i++ { //n==4,n==5,n==6
//
//		if n-i == 1 {
//			break
//		} else {
//			var r1, r2 []string
//			//先从map中取数据
//			if v, ok := allP[i]; ok {
//				r1 = v
//			} else {
//				r1 = g(i)
//			}
//			if v, ok := allP[n-i]; ok {
//				r2 = v
//			} else {
//				r2 = g(n - i)
//			}
//
//			for _, v1 := range r1 {
//				for _, v2 := range r2 {
//					res = append(res, v1+v2)
//				}
//			}
//		}
//	}
//	for k, v := range res {
//		for j := k + 1; j < len(res); j++ {
//			if v == res[j] {
//				arr := append(res[:j], res[j+1:]...)
//				res = arr
//			}
//		}
//	}
//	allP[n] = res
//	return res
//}
//
//func generateP(n int) []string {
//
//	result := make([]string, 0)
//	if n == 0 {
//		return nil
//	} else if n == 1 {
//		return []string{"()"}
//	} else if n == 2 {
//		return []string{"(())", "()()"}
//	} else {
//		left, right := "(", ")"
//		res := g(n - 1)
//		length := len(res)
//		for i := 0; i < length; i++ {
//			result = append(result, left+res[i]+right)
//			result = append(result, left+right+res[i])
//			result = append(result, res[i]+left+right)
//		}
//	}
//
//	//去重
//	for k, v := range result {
//		for j := k + 1; j < len(result); j++ {
//			if v == result[j] {
//				arr := append(result[:j], result[j+1:]...)
//				result = arr
//			}
//		}
//	}
//	return result
//}
