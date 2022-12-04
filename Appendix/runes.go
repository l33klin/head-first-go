package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiRunes := "ABCDEF"
	utf8Runes := "你好，世界"
	fmt.Printf("len(%v): %v\n", asciiRunes, len(asciiRunes))
	fmt.Printf("len(%v): %v\n", utf8Runes, len(utf8Runes))

	fmt.Printf("utf8.RuneCountInString(%v): %v\n", asciiRunes, utf8.RuneCountInString(asciiRunes))
	fmt.Printf("utf8.RuneCountInString(%v): %v\n", utf8Runes, utf8.RuneCountInString(utf8Runes))

	fmt.Printf("slice[3:] of %v: %v\n", asciiRunes, asciiRunes[3:])
	fmt.Printf("slice[3:] of %v: %v\n", utf8Runes, utf8Runes[3:])

	asciiBytes := []byte(asciiRunes)
	utf8Bytes := []byte(utf8Runes)

	fmt.Printf("[]byte of %v: %v\n", asciiRunes, asciiBytes)
	fmt.Printf("[]byte of %v: %v\n", utf8Runes, utf8Bytes)

	fmt.Printf("convert slice[4:] of []byte(%v) to string: %v\n", asciiRunes, string(asciiBytes[4:]))
	fmt.Printf("convert slice[4:] of []byte(%v) to string: %v\n", utf8Runes, string(utf8Bytes[4:]))

}
