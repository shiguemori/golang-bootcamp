package main

import "fmt"

func main() {
	n := 10
	switch {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	}
}