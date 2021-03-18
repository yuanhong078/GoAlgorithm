package main

import "fmt"

func main() {
	var s string
	s="iii小呀国"

	s1:=s[2:8]
	fmt.Println(s1, len(s1))
	fmt.Println()
	for _,v:=range s{	//打印utf-8编码
		fmt.Println(v)
	}
	fmt.Println()
	for _,v:=range s{	//打印每个字符的值
		fmt.Printf("%c",v)	//iii小呀国
	}
	fmt.Println()
	s2:=" 世界"
	s3:=s2[0:]
	bs:=[]byte(s2)
	bs[2]=','
	fmt.Println(s3)
	s2=" 中国"	//另外分配内存
	fmt.Print(s3)
}


