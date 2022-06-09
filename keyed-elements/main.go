package main

import "fmt"

func main() {
	names := [...]string{
		4: "Andrei",
	}
	fmt.Printf("%q\n", names[0])
	fmt.Printf("%d\n", len(names))
}
