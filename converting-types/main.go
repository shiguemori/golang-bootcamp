package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x = 3
	var y = 3.2

	x = x * int(y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	x = int(float64(x) * y)
	fmt.Println(x)

	var a = 5
	var b int64 = 2

	a = int(b)

	_ = a

	s := string(99)
	fmt.Println(s)
	fmt.Println(string(34234))

	var myStr = fmt.Sprintf("%f", 5.12)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	var result, err = strconv.ParseFloat("3.142", 64)
	if err == nil {
		fmt.Printf("Type: %T, Value: %v\n", result, result)
	} else {
		fmt.Println("Cannot convert to float64!")
	}

	i, err := strconv.Atoi("-50")
	s = strconv.Itoa(20)
	fmt.Printf("i Type is %T, i value is %v\n", i, i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)
}
