package main

import (
	"fmt"
)

func main() {
	// boolean
	n := true
	m := false
	// var n bool = true
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
	// Integer
	c := 42
	fmt.Printf("%v, %T\n", c, c)
	var r uint16 = 44
	fmt.Printf("%v, %T\n", r, r)
	// bit operation
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  //0010
	fmt.Println(a | b)  //1011
	fmt.Println(a ^ b)  // XOR: Sets each bit to 1 if only one of two bits is 1, 1001
	fmt.Println(a &^ b) // 0100, neither have the bit set

	d := 8
	fmt.Println(d << 3)
	fmt.Println(d >> 3)

	// complex numebr
	var complex_a complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(complex_a), real(complex_a))
	fmt.Printf("%v, %T\n", imag(complex_a), imag(complex_a))

	var complex_b complex128 = 5 + 12i
	fmt.Printf("%v, %T\n", complex_b, complex_b)

	// Text types
	// string, any UTF-8 char, immutable
	// Can be concatenated with plus(+) operator
	// Can be converted to []byte
	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)
	s2 := "This is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
	q := []byte(s)
	fmt.Printf("%v, %T\n", q, q)
	// [84 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8 <- byte

	// Rune
	// UTF-32, alias for int32
	// Special methods normally required to process, e.g. strings.Reader# ReadRune
	character := 'A' // rune (characters in Go are represented using `rune` data type)
	asciiValue := int(character)
	fmt.Printf("Ascii Value of %c = %d\n", character, asciiValue)
	var ru rune = 'a'
	fmt.Printf("%v, %T\n", ru, ru) // 97, int32
}
