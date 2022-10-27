package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		var k, n int
		fmt.Scan(&k, &n)
		fmt.Println((k/n)*(k/n))
	}
}
