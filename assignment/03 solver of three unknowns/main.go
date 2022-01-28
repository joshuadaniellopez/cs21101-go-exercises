package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/solve", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	coefficient := r.FormValue("coef")

	var a1, b1, c1, d1, a2, b2, c2, d2, a3, b3, c3, d3 int

	var statement string

	n, err := fmt.Sscanf(coefficient, "%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d", &a1, &b1, &c1, &d1, &a2, &b2, &c2, &d2, &a3, &b3, &c3, &d3)
	if err != nil {
		log.Fatal(err)
	}

	if n == 12 {
		statement = fmt.Sprintf("System: \n")
		statement = fmt.Sprintf("%s%dx + %dy + %dz = %d\n", statement, a1, b1, c1, d1)
		statement = fmt.Sprintf("%s%dx + %dy + %dz = %d\n", statement, a2, b2, c2, d2)
		statement = fmt.Sprintf("%s%dx + %dy + %dz = %d\n", statement, a3, b3, c3, d3)
	}

	var determinant, determinant_x, determinant_y, determinant_z int

	determinant = (a1 * ((b2 * c3) - (c2 * b3))) - (b1 * ((a2 * c3) - (c2 * a3))) + (c1 * ((a2 * b3) - (b2 * a3)))
	determinant_x = (d1 * ((b2 * c3) - (c2 * b3))) - (b1 * ((d2 * c3) - (c2 * d3))) + (c1 * ((d2 * b3) - (b2 * d3)))
	determinant_y = (a1 * ((d2 * c3) - (c2 * d3))) - (d1 * ((a2 * c3) - (c2 * a3))) + (c1 * ((a2 * d3) - (d2 * a3)))
	determinant_z = (a1 * ((b2 * d3) - (d2 * b3))) - (b1 * ((a2 * d3) - (d2 * a3))) + (d1 * ((a2 * b3) - (b2 * a3)))

	fmt.Printf(statement)
	fmt.Printf("D: %d\n", determinant)
	fmt.Printf("Dx: %d\n", determinant_x)
	fmt.Printf("Dy: %d\n", determinant_y)
	fmt.Printf("Dz: %d\n", determinant_z)

	if determinant == 0 {
		if determinant_x == determinant_y && determinant_y == determinant_z && determinant_z == 0 {
			statement = fmt.Sprintf("%s\ndependent - multiple solutions\n", statement)
			fmt.Printf("dependent - multiple solutions\n")
		} else {
			statement = fmt.Sprintf("%s\nincosistent - no solution\n", statement)
			fmt.Printf("incosistent - no solution\n")
		}
	} else {
		var x, y, z float32
		x = float32(determinant_x) / float32(determinant)
		y = float32(determinant_y) / float32(determinant)
		z = float32(determinant_z) / float32(determinant)

		statement = fmt.Sprintf("%s\nx = %.2f, y = %.2f, z = %.2f\n", statement, x, y, z)
		fmt.Printf("x = %.2f, y = %.2f, z = %.2f\n", x, y, z)
	}

	fmt.Fprintf(w, "%s", statement)
}
