package main

import "fmt"

func main()  {
	var t,a,b int8
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(a+b)
	}
}
