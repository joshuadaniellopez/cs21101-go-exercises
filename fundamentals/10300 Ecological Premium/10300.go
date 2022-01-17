package main

import (
	"fmt"
)

func main() {
	var n, f, farmyard, animalsCount, environmentScore int

	fmt.Scan(&n)

	if n >= 20 && n <= 0 {
		fmt.Println("Invalid # of test cases.")
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&f)

		var summativeBurden int = 0

		if f >= 20 && f <= 0 {
			fmt.Println("Invalid # of farmers.")
		}

		for j := 0; j < f; j++ {
			fmt.Scan(&farmyard, &animalsCount, &environmentScore)

			var premium float32 = float32((float32(farmyard) / float32(animalsCount)) * float32(environmentScore))
			summativeBurden = summativeBurden + int(float32(premium)*float32(animalsCount))
		}
		fmt.Println(summativeBurden)
	}
}
