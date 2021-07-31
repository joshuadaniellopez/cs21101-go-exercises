package main

import "fmt"

func main() {
	var T, a, b, c, i uint32

	fmt.Scan(&T)

	if T > 20 {
		println("Invalid # of test cases.")
	}

	for i = 0; i < T; i++ {
		fmt.Scan(&a, &b, &c) 

		if a < 1000 && a > 10000 || b < 1000 && b > 10000 || c < 1000 && c > 10000 {
			print("Invalid integer input.")
		}

		// Sort
		bubble_sort(&a, &b, &c)

		// Get the middle value
		fmt.Printf("Case %v: %v\n", i + 1, b)
	}

}

func bubble_sort(a *uint32, b *uint32, c *uint32) {
	var temp uint32
	if *b < *a {
		temp = *a 
		*a = *b 
		*b = temp
	} 

	if *c < *b  {
		temp = *b 
		*b = *c 
		*c = temp

		if *b < *a {
			temp := *a 
			*a = *b 
			*b = temp
		} 
	}
}