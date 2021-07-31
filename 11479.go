package main

import (
	"fmt"
)

func main() {
	var T, a, b, c, i uint32

	fmt.Scan(&T) 

	if T > 20 {
		println("Invalid # of test cases.")
	}

	for i = 0; i < T; i++ {
		fmt.Scan(&a, &b, &c) 

		
		if a + b <= c || a + c <= b || b + c <= a { // Check if triangle is valid
			println("Case 1: Invalid")
		} else if a == b && b == c { // Check if equilateral
			println("Case 2: Equilateral")
		} else if a == b || b == c || c == a { // Check if isoceles
			println("Case 3: Isoceles")
		} else { // Remaining case
			println("Case 4: Scalene")
		}
	}
}