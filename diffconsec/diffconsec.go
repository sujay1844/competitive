package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	for ; t>0 ; t-- {
		var n int
		fmt.Scanln(&n)
		var s string
		fmt.Scanln(&s)
		chars := []rune(s)
		dups := 0
		for i:=0 ; i<n ; i++ {
			if i!=0 {
				if chars[i] == chars[i-1] {
					dups++
				}
			}
		}
		fmt.Println(dups)
	}
}
