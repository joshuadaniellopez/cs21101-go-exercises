package main

import (
	"fmt"
)

func main() {
	var v, t int32
	_, err := fmt.Scan(&v, &t)

	for err == nil {
		println(v * t * 2)
		_, err = fmt.Scan(&v, &t)
	}
}
