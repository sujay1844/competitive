package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		max_count := 0
		pointers := make([]interface{}, n)
		for i := range arr {
			pointers[i] = &arr[i]
		}
		n, _ = fmt.Scanln(pointers...)
		arr = arr[:n]
		for idx1, i := range arr {
			count := 0
			for idx2, j := range arr {
				if i == j && idx1 != idx2 {
					count++
				}
			}
			if count > max_count {
				max_count = count
			}
		}
		fmt.Println(n-max_count-1)
	}
}
