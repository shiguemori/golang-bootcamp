package main

import "fmt"

type age int
type oldAge age
type veryOldAge age

func main() {

	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint

	x = uint(s1)
	_ = x

	var s3 = speed(x)
	_ = s3
}
