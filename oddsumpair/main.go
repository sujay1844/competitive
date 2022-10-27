package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	for ; t>0 ; t-- {
		var a, b, c int
		fmt.Scanln(&a, &b, &c)
		if ((a+b)%2 != 0) || ((b+c)%2 != 0) || ((c+a)%2 != 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
