package main

import (
	"fmt"
)

func main() {

	outer := 19
	_ = outer

	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

outer:

	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
				break outer
			}
		}
	}

	fmt.Println("Next instruction after the break.")

	i := 0
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

}
