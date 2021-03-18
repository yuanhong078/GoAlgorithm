package main

import "fmt"

func main() {
	arr:=[]int{1,2,3,4,5,6}
	fmt.Println(binarySearch(arr,3))
}

func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (r + l) / 2
		if target == arr[mid] {
			return mid
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
