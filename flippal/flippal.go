package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		var n int
		fmt.Scan(&n)
		var str string
		fmt.Scan(&str)
		chars := []rune(str)
		ones := 0
		zeroes := 0
		for _, i := range chars {
			if string(i) == "1" {
				ones++
			}
			if string(i) == "0" {
				zeroes++
			}
		}
		if ones % 2 != 0 && zeroes % 2 != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
