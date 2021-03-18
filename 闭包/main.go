package main

import "fmt"

func main() {
	fs:=create()
	for i := 0;i<len(fs);i++{
		fs[i]()
	}
}

func create() (fs [2]func()){
	for i:=0;i<2;i++{
		fs[i]=func(){
			fmt.Println(i)
		}
	}
	return
}