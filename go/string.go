package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func checkLength() {
	// Ʃ (https://unicode-table.com/en/01A9/)
	// This single character is encoded using at least two bytes
	l := len("Ʃ")
	if l == 1 {
		fmt.Printf("👍 Unicode character length\n")
	} else {
		fmt.Printf("👎 Unicode character length. Got: %d\n", l)
		// The solution is to use the uf8 package
		l = utf8.RuneCountInString("Ʃ")
		if l == 1 {
			fmt.Printf("👍 Unicode character length using utf8 package\n")
		} else {
			fmt.Printf("👎 Unicode character length using utf8 package. Got: %d\n", l)
		}
	}
}

func checkComparison() {
	sequence1 := "Voil\u00E0"       // à
	sequence2 := "Voil\u0061\u0300" // a + combining grave accent

	if sequence1 == sequence2 {
		fmt.Println("👍 Unicode string comparison using ==")
	} else {
		fmt.Println("👎 Unicode string comparison using ==")
	}

	// https://blog.golang.org/normalization

	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	norm1, _, _ := transform.String(t, sequence1)
	norm2, _, _ := transform.String(t, sequence2)
	if strings.EqualFold(norm1, norm2) {
		fmt.Println("👍 Unicode string comparison using golang.org/x/text package")
	} else {
		fmt.Println("👎 Unicode string comparison using golang.org/x/text package")
	}
}

func checkCaseFolding() {
	sequence1 := "Voil\u00E0" // à
	sequence2 := "Voil\u00C0" // À

	res := strings.ToUpper(sequence1)
	if res != sequence2 {
		fmt.Println("👍 Unicode case folding using strings.ToUpper()")
	} else {
		fmt.Println("👎 Unicode case folding using strings.ToUpper()")
	}

	res = strings.ToLower(sequence2)
	if res != sequence1 {
		fmt.Println("👍 Unicode case folding using strings.ToLower()")
	} else {
		fmt.Println("👎 Unicode case folding using strings.ToLower()")
	}

	if strings.EqualFold(sequence1, sequence2) {
		fmt.Println("👍 Unicode case folding using strings.EqualFold()")
	} else {
		fmt.Println("👎 Unicode case folding using strings.EqualFold()")
	}
}

func checkRegex() {
	// µ https://unicode-table.com/en/00B5/
	s := "100 µAh 10 mAh"
	ampHoursValues := regexp.MustCompile(`\d+ \wAh`)
	res := ampHoursValues.FindAllString(s, -1)
	if len(res) == 2 {
		fmt.Println("👍 Unicode regex using \\w")
	} else {
		fmt.Println("👎 Unicode regex using \\w")
	}

	// In Go, the \w shorthand class only matches ASCII letters.
	// We need a Unicode character class \p{L}.
	// L (letter), Ll (lowercase), Lm (modifier), M (Mark), N (Number),
	// P (Punctuation), S (Symbol), ...

	ampHoursValues = regexp.MustCompile(`\d+ \p{L}Ah`)
	res = ampHoursValues.FindAllString(s, -1)
	if len(res) == 2 {
		fmt.Println("👍 Unicode regex using \\p{L}")
	} else {
		fmt.Println("👎 Unicode regex using \\p{L}")
	}

}

func main() {
	checkLength()
	checkComparison()
	checkCaseFolding()
	checkRegex()
}

// Output:
// ------
// 1. Length
// 👎 Unicode character length. Got: 2
// 👍 Unicode character length using utf8 package
// 2. Text Comparison
// 👎 Unicode string comparison using ==
// 👍 Unicode string comparison using golang.org/x/text package
// 3. Text Case Folding
// 👍 Unicode case folding using strings.ToUpper()
// 👍 Unicode case folding using strings.ToLower()
// 👍 Unicode case folding using strings.EqualFold()
// 4. Regex
// 👎 Unicode regex using \w
// 👍 Unicode regex using \p{L}
