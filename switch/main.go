package main

import "fmt"

func main() {

	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python! You don't use { } but indentation !! ")
	case "Go", "golang":
		fmt.Println("Good, Go for Go!. You are using {}!")
	default:
		fmt.Println("Any other programming language is a good start!")
	}

	n := 5

	switch true {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	default:
		fmt.Println("Never here!")
	}

	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
