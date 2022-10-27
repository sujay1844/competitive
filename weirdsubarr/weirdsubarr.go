package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Scanln(&t)
	for ; t>0 ; t-- {
		var n int
		fmt.Scanln(&n)
		s, _ := in.ReadString('\n')
		s = s[:len(s)-1]
		res := strings.Split(s, " ")
		a := make([]int, len(res))
		for i:=0 ; i<len(a) ; i++ {
			a[i], _ = strconv.Atoi(res[i])
		}
		counter := 0
		for i:=0 ; i<len(a) ; i++ {
			for j:=i ; j<len(a) ; j++ {
				if canBeWeird(a[i:j+1]) {
					counter++
				}
			}
		}
		fmt.Println(counter)
	}
}

func isWeird(arr []int) bool {
	for i:=1; i<len(arr); i++ {
		if arr[i] >= arr[i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func canBeWeird(arr []int) bool {
	for i:=0; i<len(arr); i++ {
		if isWeird(arr) {
			return true
		} else {
			arr[i] = -arr[i]
		}
	}
	return false
}
