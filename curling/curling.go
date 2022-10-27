// Google KickStart Round G 2022
// https://codingcompetitions.withgoogle.com/kickstart/round/00000000008cb2e1/0000000000c17c82
package main

import "fmt"

func main() {
	var T int
    fmt.Scanln(&T)
    for k:=1 ; k<=T ; k++ {
		var Rs, Rh, N, M int
		score1, score2 := 0, 0
        fmt.Scanln(&Rs, &Rh)
        fmt.Scanln(&N)
		var x, y int
		distMin1, distMin2 := 400040001, 400040001
		dist1 := make([]int, N)
        for i:=0 ; i<N ; i++ {
            fmt.Scanln(&x, &y)
			dist1[i] = distance(x, y)
            if i > 0 {
                if dist1[i] < dist1[i-1] {
					distMin1 = dist1[i]
                }
            } else {
				distMin1 = dist1[i]
			}
        }
        fmt.Scanln(&M)
		dist2 := make([]int, M)
		var z, w int
        for j:=0 ; j<M ; j++ {
            fmt.Scanln(&z, &w)
			dist2[j] = distance(z, w)
            if j > 0 {
                if dist2[j] < dist2[j-1] {
					distMin2 = dist2[j]
                }
            } else {
				distMin2 = dist2[j]
			}
        }
        if distMin1 < distMin2 {
            for i:=0 ; i<N ; i++ {
                if dist1[i] < distMin2 && dist1[i] <= (Rs+Rh)*(Rs+Rh){
                    score1++
                }
            }
        } else {
            for j:=0 ; j<M ; j++ {
                if dist2[j] < distMin1 && dist2[j] <= (Rs+Rh)*(Rs+Rh){
                    score2++
                }
            }
        }
        fmt.Printf("Case #%d: %d %d\n", k, score1, score2)
    }
}

func distance(x, y int) int {
    return (x*x) + (y*y)
}
