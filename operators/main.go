package main

import "fmt"

func main() {
	a, b := 10, 5.5

	fmt.Println(a + 5)
	fmt.Println(3.1 - b)
	fmt.Println(a * a)
	fmt.Println(a / a)
	fmt.Println(11 / 5)

	fmt.Println(a * int(b))
	fmt.Println(float64(a) * b)

	x := 10
	x++
	x--

	a = 10
	a += 2
	a -= 3
	a *= 2
	a /= 3
	a %= 5

	fmt.Println(5 == 6)
	fmt.Println(5 != 6)
	fmt.Println(10 > 10)
	fmt.Println(10 >= 10)
	fmt.Println(5 < 5)
	fmt.Println(5 <= 5)

	fmt.Println(0 < 2 && 4 > 1)
	fmt.Println(1 > 5 || 4 > 5)
	fmt.Println(!(1 > 2))

}
