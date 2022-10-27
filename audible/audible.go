package main

import "fmt"

func main() {
	var T, X int
	
	fmt.Scan(&T)

	for ; T > 0 ; T-- {
		fmt.Scan(&X)
		if X >= 67 && X <= 45000 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
