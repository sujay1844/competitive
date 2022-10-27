package main

import "fmt"

func main()  {
	var t, x int8
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		fmt.Scan(&x)
		fmt.Println(2*x)
	}
}
