package main

import (
	"fmt"
)

func main() {
	var t, n, m int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n, &m)
		fmt.Println((n / 3) * (m / 3))
	}
}
