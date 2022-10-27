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
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		var n int
		fmt.Scan(&n)
		s, _ := in.ReadString('\n')
		s = s[:len(s)-1]
		res := strings.Split(s, " ")
		a := make([]int, len(res))
		for i:=0 ; i<len(a) ; i++ {
			a[i], _ = strconv.Atoi(res[i])
		}
		output := ""
		partitions := 0
		for i:=0 ; i<len(a) ; i++ {
			for j:=i ; j<len(a) ; j++ {
				fmt.Println(a[i:j])
				if isPal(a[i:j]) {
					output += strconv.Itoa(j-i) + " "
					partitions++
				}
			}
		}
		if output == "" {
			fmt.Println(-1)
		} else {
			fmt.Println(partitions)
			fmt.Println(output)
		}

	}
}

func isPal(arr []int) bool {
	for i:=0 ; i<len(arr)/2 ; i++ {
		if arr[i] != arr[len(arr)-i] {
			return false
		}
	}
	return true
} 
