package main

import "fmt"

func main() {

	var cities []string

	fmt.Println("cities is equal to nil: ", cities == nil)

	fmt.Printf("cities: %#v\n", cities)

	numbers := []int{2, 3, 4, 5}

	fmt.Println(numbers)

	nums := make([]int, 2)

	fmt.Println(nums)

	type names []string

	friends := names{"Dan", "Maria"}

	fmt.Println(friends)

	x := numbers[0]

	fmt.Println("x is", x)

	numbers[1] = 200

	fmt.Printf("%#v\n", numbers)

	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers[i])
	}

	var n []int
	n = numbers
	_ = n
	var nn []int
	fmt.Println(nn == nil)
	mm := []int{}
	fmt.Println(mm == nil)
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq = true
	if len(a) != len(b) {
		eq = false
	}

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

}
