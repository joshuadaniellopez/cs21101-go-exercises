package main

import (
	"fmt"
)

func main () {
	var i uint32

	for i = 0; i < 2000; i++ {
		var word string
		fmt.Scanln(&word)

		if word == "HELLO" {
			fmt.Printf("Case %v: ENGLISH\n", i + 1)
		} else if word == "HOLA" {
			fmt.Printf("Case %v: SPANISH\n", i + 1)
		} else if word == "HALLO" {
			fmt.Printf("Case %v: GERMAN\n", i + 1)
		} else if word == "BONJOUR" {
			fmt.Printf("Case %v: FRENCH\n", i + 1)
		} else if word == "CIAO" {
			fmt.Printf("Case %v: ITALIAN\n", i + 1)
		} else if word == "ZDRAVSTVUJTE" {
			fmt.Printf("Case %v: RUSSIAN\n", i + 1)
		} else {
			fmt.Printf("Case %v: UNKNOWN\n", i + 1)
		}
	}
}