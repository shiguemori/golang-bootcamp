package main

import "fmt"

func main() {

	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	var i3 int64 = -324_567_345
	fmt.Printf("%T\n", i3)
	fmt.Printf("i3 is %d\n", i3)

	var f1, f2, f3 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	var r = 'f'
	fmt.Printf("%T\n", r)
	fmt.Printf("%x\n", r)
	fmt.Printf("%c\n", r)

	var b = true
	fmt.Printf("%T\n", b)

	var s = "Hello Go!"
	fmt.Printf("%T\n", s)

	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	var cities = []string{"London", "Bucharest", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}
	fmt.Printf("%T\n", balances)

	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	var x = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr)

	fmt.Printf("%T\n", f)
}

func f() {
}
