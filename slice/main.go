package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//s1 := make([]int, 4, 10)
	//s2 := s1
	//s2[0] = 100
	//
	//s2 = append(s2, 1, 2, 3, 4)
	//s1 = s1[0:5]
	//println("s11: ", s1, "s12: ", s2) //s1==s2
	//fmt.Println("s11: ", s1, cap(s1))
	//fmt.Println("s12: ", s2, cap(s2))
	//time.Sleep(5 * time.Second)
	//
	//s1 = append(s1, s2...)
	//println("s21: ", s1, "s22: ", s2) //s1!=s2
	//fmt.Println("s21: ", s1, cap(s1)) //20
	//fmt.Println("s22: ", s2, cap(s2)) //10
	//time.Sleep(5 * time.Second)
	//
	//s2 = append(s2, 1000)             //分配新的内存
	//println("s31: ", s1, "s32: ", s2) //s1!=s2
	//fmt.Println("s31: ", s1, cap(s1)) //10
	//fmt.Println("s32: ", s2, cap(s2)) //20

	//s := []int{1, 1, 1}

	s:=make([]int,3,5)
	fmt.Println(unsafe.Pointer(&s))
	fmt.Println(unsafe.Pointer(&s[1]))
	f(s)
	fmt.Println(s)

}

func f(a []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
	      i++
	  }
	*/
	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(&a[1]))
	for i := range a {
		a[i] += 2

	}

	a=append(a,100)
	fmt.Println("===")
	s:=fmt.Sprintf("%q",a)
	fmt.Println(s)
	fmt.Println("===")
	fmt.Println(a)
}