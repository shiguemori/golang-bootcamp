package main

import "fmt"

func main() {
	const days int = 7
	const pi = 3.14
	fmt.Printf("%d\n", days)
	fmt.Printf("%f\n", pi)

	const (
		a         = 5
		b float64 = 0.1
	)
	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", b)

	const n, m int = 4, 5
	fmt.Printf("%d\n", n)
	fmt.Printf("%d\n", m)

	const (
		min1 = -500
		max1
		max2
	)
	fmt.Printf("%d\n", min1)
	fmt.Printf("%d\n", max1)
	fmt.Printf("%d\n", max2)

	const temp int = 100
	fmt.Printf("%d\n", temp)

	t := 5
	fmt.Printf("%d\n", t)

	const l1 = len("Hello")
	fmt.Printf("%d\n", l1)

	str := "Hello"
	fmt.Printf("%d\n", len(str))

	_, _ = t, str

	const x = 5
	const y float64 = 1.1
	fmt.Printf("%d\n", x)
	fmt.Printf("%f\n", y)

	var v1 = 5
	var v2 = 1.1
	fmt.Printf("%d\n", v1)
	fmt.Printf("%f\n", v2)

	fmt.Println(x * y)

	_, _ = v1, v2

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		North = iota
		East
		South
		West
	)
	fmt.Println(North, East, South, West)
	const (
		c11 = iota * 2
		c22
		c33
	)
	fmt.Println(c11, c22, c33)
}
