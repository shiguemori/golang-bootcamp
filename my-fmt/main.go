package main

import "fmt"

func main() {

	fmt.Println("Hello Go World!")

	var name, age = "Andrei", 35
	fmt.Println(name, "is", age, "years old.")

	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> double-quoted string
	// %v -> value (any)
	// %#v -> a Go-syntax representation of the value
	// %T -> value Type
	// %t -> bool (true or false)
	// %p -> pointer (address in base 16, with leading 0x)
	// %c -> char (rune) represented by the corresponding Unicode code point

	a, b, c := 10, 15.5, "Gophers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c)
	fmt.Printf("%q\n", c)
	fmt.Printf("%v\n", grades)
	fmt.Printf("%#v\n", grades)
	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
	fmt.Printf("The address of a: %p\n", &a)
	fmt.Printf("%c and %c\n", 100, 51011)

	const pi float64 = 3.14159265359
	fmt.Printf("pi is %.4f\n", pi)

	fmt.Printf("255 in base 2 is %b\n", 255)
	fmt.Printf("101 in base 16 is %x\n", 101)

	s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
	fmt.Println(s)
}
