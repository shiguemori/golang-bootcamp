package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var1 := 'a'
	fmt.Printf("Type: %T, Value:%d \n", var1, var1)

	str := "ţară"

	fmt.Println(len(str))

	m := utf8.RuneCountInString(str)
	fmt.Println(m)

	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])

		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 10))

	for i, r := range str {
		fmt.Printf("%d -> %c", i, r)
	}

}
