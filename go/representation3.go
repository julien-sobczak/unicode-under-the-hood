package main

import (
	"fmt"
)

func main() {

	s := "Hey 🤚!"

	for index, runeValue := range s {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
}
// Output:
// U+0048 'H' starts at byte position 0
// U+0065 'e' starts at byte position 1
// U+0079 'y' starts at byte position 2
// U+0020 ' ' starts at byte position 3
// U+1F91A '🤚' starts at byte position 4
// U+0021 '!' starts at byte position 8
