package main

import (
	"fmt"
)

func main() {
	var a, b, c uint32

	_, err := fmt.Scan(&a, &b, &c)

	for err == nil {

		bubble_sort(&a, &b, &c)
		if a == 0 && a == b && b == c {
			return
		}

		if a*a+b*b == c*c {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
		_, err = fmt.Scan(&a, &b, &c)
	}

}

func bubble_sort(a *uint32, b *uint32, c *uint32) {
	var temp uint32
	if *b < *a {
		temp = *a
		*a = *b
		*b = temp
	}

	if *c < *b {
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
