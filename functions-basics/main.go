package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1() function")
}

func f2(a int, b int) {

	fmt.Println("Sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)

}

func f5(a, b int) (int, int) {
	return a * b, a + b
}

func sum(a, b int) (s int) {
	fmt.Println("s:", s)
	s = a + b

	return
}

func main() {

	f1()
	f3(3, 4, 5, 4., 5.5, "ss")
	fmt.Println(f4(5.1))
	a, b := f5(6, 7)
	fmt.Printf("a:%d, b:%d\n", a, b)
	ss := sum(4, 5)
	fmt.Println(ss)

}
