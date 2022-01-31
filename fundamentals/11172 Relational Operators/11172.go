package main

import (
	"fmt"
)

func main(){
	var t, a, b int
	
	fmt.Scan(&t) 

	if t >= 15 {
		println("Invalid # of test cases.")
	}

	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b) 
		
		if a > b {
			println(">")
		} else if a == b {
			println("=")
		} else {
			println("<")
		}
	}
}