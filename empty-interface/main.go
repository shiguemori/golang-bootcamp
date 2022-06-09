package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {

	var empty interface{}

	empty = 5
	fmt.Println(empty)
	empty = "Go"
	fmt.Println(empty)

	empty = []int{2, 34, 4}
	fmt.Println(empty)

	fmt.Println(len(empty.([]int)))

	you := person{}

	you.info = "You name"
	you.info = 40
	you.info = []float64{4.5, 6., 8.1}

	fmt.Println(you.info)
}
