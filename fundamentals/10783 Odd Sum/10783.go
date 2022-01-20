package main

import "fmt"

func main() {
	var t int       // Test cases
	var a, b uint32 // Input Integers
	var odd_sum uint32

	_, err := fmt.Scan(&t) // Retrieve Test Case
	if err != nil {
		return
	}

	for c := 0; c < t; c++ {
		fmt.Scan(&a, &b)

		odd_sum = 0
		for i := a; i <= b; i++ {
			if i%2 == 0 {
				continue
			} else {
				odd_sum += i
			}
			i++
		}
		fmt.Printf("Case %v: ", c+1)
		fmt.Println(odd_sum)
	}
}
