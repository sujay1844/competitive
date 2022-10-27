package main

import "fmt"

func main()  {
	var x, y, z int8
	var t int16
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		fmt.Scanf("%d %d %d", &x, &y, &z)
		if x>y && x>z {
			fmt.Println("Setter")
		} else if y>x && y>z {
			fmt.Println("Tester")
		} else if z>x && z>y {
			fmt.Println("Editorialist")
		}
	}
}
