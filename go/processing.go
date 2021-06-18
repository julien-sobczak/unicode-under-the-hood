package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	fmt.Println("à" == "à")
	// Output: false

	fmt.Println("\u00E0 == \u0061\u0300")
	fmt.Println("\u00E0" == "\u0061\u0300")

	norm1 := norm.NFD.String("\u00E0")
	norm2 := norm.NFD.String("\u0061\u0300")
	fmt.Println(norm1 == norm2)
	// Output: true
}
