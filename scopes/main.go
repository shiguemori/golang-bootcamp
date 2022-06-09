package main

import (
	"fmt"

	f "fmt"
)

const done = false

func main() {

	var task = "Running:"
	fmt.Println(task, done)
	f.Println("Bye bye!")

	const done = true
	fmt.Printf("done in main(): %v\n", done)
	f1()
}

func f1() {
	fmt.Printf("done in f(): %v\n", done)
}
