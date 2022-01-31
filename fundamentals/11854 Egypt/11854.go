package main

import (
	"fmt"
)

func main() {
	var side_a, side_b, side_c, a, b, c uint32

	_, err := fmt.Scan(&a, &b, &c)

	for err == nil {
		if a >= b && a >= c {
			side_c = a
			side_a = b
			side_b = c
		} else if b >= a && b >= c {
			side_c = b
			side_a = a
			side_b = c
		} else {
			side_c = c
			side_a = a
			side_b = b
		}

		if side_a*side_a+side_b*side_b == side_c*side_c {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
		_, err = fmt.Scan(&a, &b, &c)
	}
}
