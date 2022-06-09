package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
		fmt.Printf("i = %v, j = %v\n", i, j)
	}

	// sum := 0
	// for {
	//  sum++
	// }
	// fmt.Println(sum)
}
