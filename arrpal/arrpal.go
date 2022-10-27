package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		var n int
		fmt.Scan(&n)
		arr := Scann(n)
		counter := 0
		possible := true
		for i:=len(arr)/2-1 ; i>=0 ; i-- {
			diff := arr[len(arr)-i-1] - arr[i]
			if diff >= 0 {
				for j:=0 ; j<=i ; j++ {
					arr[j] += diff
				}
				counter += diff
			} else {
				fmt.Println("-1")
				possible = false
				break
			}
		}
		if possible {
			fmt.Println(counter)
		}
	}
}
func Scann(n int) []int {
	arr := make([]int, n)
	pointers := make([]interface{}, n)
	for i := range arr {
		pointers[i] = &arr[i]
	}
	n, _ = fmt.Scanln(pointers...)
	arr = arr[:n]
	return arr
}
