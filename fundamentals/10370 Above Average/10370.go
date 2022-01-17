package main

import (
	"fmt"
	"sort"
)

func main() {
	var C, N int

	fmt.Scan(&C)

	for i := 0; i < C; i++ {
		fmt.Scan(&N)

		var grades []int
		var averageGrade float32

		for j := 0; j < N; j++ {
			var grade int
			fmt.Scan(&grade)
			grades = append(grades, grade)
			averageGrade = averageGrade + float32(grade)
		}

		averageGrade = averageGrade / float32(N)
		sort.Ints(grades)

		var countPassed int = 0
		for j := N - 1; j >= 0; j-- {
			if float32(grades[j]) > averageGrade {
				countPassed = countPassed + 1
			}
		}

		var passersPercentage float32 = float32(countPassed) / float32(N) * 100
		fmt.Printf("%3.3f%% \n", passersPercentage)

	}
}
