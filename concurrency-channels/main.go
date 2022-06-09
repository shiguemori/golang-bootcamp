package main

import "fmt"

func main() {

	var c1 chan int
	fmt.Println(c1)

	c1 = make(chan int)
	fmt.Println(c1)

	c2 := make(chan int)
	_ = c2

	c3 := make(<-chan string)
	c4 := make(chan<- string)
	fmt.Printf("%T, %T, %T\n", c1, c3, c4)
	c1 <- 10
	num := <-c1
	_ = num

	fmt.Println(<-c1)
	close(c1)

}
