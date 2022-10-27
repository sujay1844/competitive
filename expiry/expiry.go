package main

import "fmt"

func main()  {
	var t int16
	var n, m, k int8
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		fmt.Scanf("%d %d %d", &n, &m, &k)
		if m*k >= n {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
