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
			println("Invalid integer input.")
		}

		if a >= b && a <= c {
			fmt.Printf("Case %v: %v\n", i + 1, a)
		} else if b >= a && b <= c || b <= a &&  b >= c{
			fmt.Printf("Case %v: %v\n", i + 1, b)
		} else {
			fmt.Printf("Case %v: %v\n", i + 1, c)
		}
	}

}