package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	for ; t>0 ; t-- {
		var x int
		fmt.Scanln(&x)
		if x >= 30 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
