package main

import (
	"fmt"
)

func main() {

	s1 := "Hi there  Go!"

	fmt.Printf("%s\n", s1)
	fmt.Printf("%q\n", s1)

	fmt.Println("He say: \"Hello!\"")

	fmt.Println(`He say: "Hello!"`)

	s2 := `Hi there Go!`
	fmt.Println(s2)

	fmt.Println("Price: 100 \nBrand: Nike")

	fmt.Println(`
Price: 100
Brand: Nike`)

	fmt.Println(`C:\Users\Andrei`)
	fmt.Println("C:\\Users\\Andrei")

	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index zero:", s3[0])

}
