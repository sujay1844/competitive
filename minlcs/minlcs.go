package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		var n int
		fmt.Scan(&n)
		var a,b string
		fmt.Scan(&a)
		fmt.Scan(&b)
		runeA := []rune(a)
		lcs := ""
		max_lcs := 0
		for i:=0 ; i<len(runeA) ; i++ {
			for j:=i ; j<len(runeA) ; j++ {
				if strings.Contains(b, lcs+string(runeA[j])) {
					lcs = lcs + string(runeA[j])
				} else {
					break
				}
			}
			if max_lcs < len(lcs) {
				max_lcs = len(lcs)
			}
			lcs = ""
		}
		fmt.Println(max_lcs)
	}
}
