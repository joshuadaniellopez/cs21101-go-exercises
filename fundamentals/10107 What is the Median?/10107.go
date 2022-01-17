package main

import (
	"fmt"
	"sort"
)

func main() {
	var inputs []int
	var input, length int

	_, err := fmt.Scan(&input)
	for err == nil {
		inputs = append(inputs, input)
		sort.Ints(inputs)

		length = len(inputs)

		if length % 2 == 0 {
			fmt.Println((inputs[length/2]) + (inputs[length/2 - 1])/2)
		} else {
			fmt.Println((inputs[length/2]))
		}
		_, err = fmt.Scan(&input)
	}
}