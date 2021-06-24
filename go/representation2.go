package main

import (
	"fmt"
)

func main() {

	s := []rune("Hey 🤚!")

	fmt.Printf("len=%d\n", len(s))
	// Print the characters
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println("")
	// Print the code points
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
// Output:
// len=6
// H e y   🤚 !
// U+0048 'H' U+0065 'e' U+0079 'y' U+0020 ' ' U+1F91A '🤚' U+0021 '!'
