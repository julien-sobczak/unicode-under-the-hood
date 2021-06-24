package main

import (
	"fmt"
)

func main() {

	s := "Hey 🤚!" // String literal stored in an UTF-8 file

	fmt.Printf("len=%d\n", len(s))
	// Print characters
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println("")
	// Print bytes
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}
}
