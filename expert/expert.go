package main

import "fmt"

func main() {
	var t int
	var x, y int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		fmt.Scan(&x, &y)
		if float64(y)/float64(x) >= 0.5  {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
