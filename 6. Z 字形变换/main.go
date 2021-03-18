package main

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

*/
func main() {

}

//0   4   8    12      3
//1 3 5 7 9  11
//2   6   10
//
//0      6      12       4
//1    5 7   11 13
//2  4   8 10   14
//3      9      15
func convert(s string, numRows int) string {
	arr := make([][]string, numRows)
	//4  6
	return arr[0][0]

}
