package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println("Len:", len(s))
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) //%x to print the byte value
	}
	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))
	fmt.Println()

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println()

	fmt.Println("Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}

}
