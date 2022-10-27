package main
import "fmt"

func main(){
	var t int
	var x int
	fmt.Scan(&t)
	for ; t>0 ; t-- {
	    fmt.Scan(&x)
	    if x > 100 {
	        fmt.Println(x-10)
	    } else {
	        fmt.Println(x)
	    }
	}
}

