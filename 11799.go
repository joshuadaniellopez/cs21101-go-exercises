package main

import (
	"fmt"
)

func main () {
	var T, N, i, j, max, temp uint32
	
	fmt.Scan(&T) 

	if T > 50 {
		println("Invalid # of test cases.")
	}

	for i = 0; i < T; i++ {
		fmt.Scan(&N)
		max = 0
		for j = 0; j < N; j++ {		
			fmt.Scan(&temp)
			if temp > max {
				max = temp
			}
		}
		
		fmt.Printf("Case %v: %v\n", i + 1, max)
	}
}