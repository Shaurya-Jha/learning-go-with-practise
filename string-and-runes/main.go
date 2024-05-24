package main

import (
	"fmt"
	"unicode/utf8"
)

// concept of character is called rune in go
// rune is an integer that represents a unicode code point

func main() {
	const s = "สวัสดี"
	fmt.Println("length: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d ", runeValue, idx)
	}

	fmt.Println("using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d ", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

// examineRune function
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("Found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
