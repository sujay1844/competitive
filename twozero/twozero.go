package main

import (
	"fmt"
	"math"
)

func main()  {
	var n, m, num_of_subsets, i, j int64
	var t int16
	fmt.Scan(&t)
	for ; t>0 ; t-- {
		fmt.Scanf("%d %d", &n, &m)
		// Looping over every possible subset of 2s and 0s
		for i=0; i<=n ; i++ {
			for j=0; j<=m ; j++ {
				// Formula for score can be simplified
				// i*2 + j*0 - (i+j) = i - j
				if (i - j) % 3 == 0 {
					if i == 0 && j == 0 {
						// Ignore empty sets
						continue
					} else if i == 0 {
						// counting duplicates
						num_of_subsets+= perm(n, i)
					} else if j == 0 {
						num_of_subsets+= perm(n, j)
					} else {
						num_of_subsets+= perm(n, i) + perm(m, j)
					}
				}
			}
		}
		num_of_subsets %= int64(math.Pow10(9)+7)
		fmt.Println(num_of_subsets)
	}
}
func perm(n, r int64) int64 {
	// nCr = n! / (n-r)! r!
	var comb int64
	comb = fact(n)/fact(r)
	return comb
}
func fact(n int64) int64 {
	var fact int64
	for fact=1; n>0 ; n-- {
		fact *= n
	}
	return fact
}
