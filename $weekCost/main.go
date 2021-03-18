package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	//res := costCount([]float64{3.6, 6.5, 8, 21, 5.6, 8, 9.9, 20, 8.36, 12, 6, 17.99, 19, 30, 5, 11, 12, 5.99, 26, 2, 5.4, 23.92, 6, 29, 59.5, 3.51, 20, 15})
	//res := costCount([]float64{7,1.99,25,8,13,6,20,23.6,243,76,23.67,23.85,7,20,2.5,9.99,94.73,11,20,9.99,82.5,25,9.9,8.9,2,12,20,22.9,4,5,11.7,3.43,6,20,4.17,4,2,14.36,62.8,263})
	//res := costCount([]float64{6,20,8,8,20,7.99,15,8.5,11.8,9.98,9.9,9,10,22,8.5,16.79,2,30,22,4.22,29.9,31.98,6,20,3.9,3.9,241,29.7,2,8.7,2.8,12.5,8.5,9.9,5.5,9,9})
	//res := costCount([]float64{7.99, 11.8, 9.9, 9.98, 9, 16.79, 2, 29.9, 31.98, 3.9, 3.9, 241, 2, 12.5, 8.7, 2.8, 8.5, 9.9, 5.5, 9.9})

	//fmt.Println(res)

	//a := []string{"2","3","4","5"}

	//z := strings.Join(a, "-")

	//c:=strings.Fields("ab ddd ab fff")
	//
	//fmt.Print(c)
	//s := make([]int, 3)
	//s[0] = 2
	//z := [1]int{2}
	//fmt.Println("1:", s, "  ", z)
	//testSlice(s, z)
	//fmt.Println("2:", s, "  ", z)

	s := "哎呦喂,123%%%"
	fmt.Println(len(s))
	r := []byte(s)
	l := utf8.RuneCount(r)
	fmt.Println(len(r), "==", l)
	e, size := utf8.DecodeRune(r)

	//fmt.Printf("%c %d", e, size)
	fmt.Println(string(e), "---", size)
	fmt.Println()
	s = ""
	for i := 0; i < 2; i++ {
		e, size := utf8.DecodeRune(r)
		r = r[size:]
		s += string(e)
		//fmt.Printf("%c %d", e, size)
	}
	fmt.Print(s)
}

func testSlice(a []int, z [1]int) {
	a[0] = 1
	a = append(a, 3)
	z[0] = 22
}

func costCount(costs []float64) float64 {
	var sum float64
	for _, v := range costs {

		sum += v
		//fmt.Println("k:",k," v: ",v," sum:",sum)
	}
	return sum
}

func change(a, b int) (x, y int) {

	c, x := a+100, 1
	d, y := b+100, 2

	e, x := 100, 3
	f, y := 200, 4
	if d > c {
		return
	}
	fmt.Println(e, f, x, y)

	time.Now().Format("2006-01-02 15:04:05.xxx")
	return //101, 102
	//return x, y  //同上
	//return y, x  //102, 101
}
