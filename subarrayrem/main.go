package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Scan(&t)
	for ; t>0; t-- {
		var n int
		fmt.Scanln(&n)
		s, _ := in.ReadString('\n')
		s = s[:len(s)-1]
		res := strings.Split(s, " ")
		a := make([]int, len(res))
		for i:=0 ; i<len(a) ; i++ {
			a[i], _ = strconv.Atoi(res[i])
		}

		score := 0
		flag := true
		for i:=0 ; i<len(a)-1 ; i++ {
			if a[i] == 0 {
				flag = false
			}
			if (a[i] == 0 && a[i+1] == 1) || (a[i] == 1 && a[i+1] == 0) {
				fmt.Println(a)
				a = append(a[:i], a[i+2:]...)
				score++
			}
		}
		if flag {
			fmt.Println(1)
		} else {
			fmt.Println(a)
			fmt.Println(score)
		}
	}
}
