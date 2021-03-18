package main

import "fmt"

type Test struct {
	name string
}

func (this *Test) Point(){
	fmt.Println(this.name)
}


func main() {

	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	for _,t := range ts {
		//fmt.Println(reflect.TypeOf(t))
		defer t.Point()
	}

}
