package main

import (
	"fmt"
)

func main() {
	var T, L, W, H, i uint32
	var flag_wontFit bool

	fmt.Scan(&T)

	for i = 0; i < T; i++ {
		fmt.Scan(&L, &W, &H)

		if L > 20 && L <= 0 {
			flag_wontFit = true
		}

		if W > 20 && L <= 0 {
			flag_wontFit = true
		}

		if H > 20 && L <= 20 {
			flag_wontFit = true
		}

		if flag_wontFit == true {
			fmt.Printf("Case %v: bad\n", i+1)
		} else {
			fmt.Printf("Case %v: good\n", i+1)
		}
	}
}
