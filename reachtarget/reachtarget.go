package main

import "fmt"

func main() {
	var t,x,y int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		fmt.Scanf("%d %d", &x, &y)
		fmt.Println(x-y)
	}
}
