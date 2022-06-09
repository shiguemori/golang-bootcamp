package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println

	result := strings.Contains("I love Go Programming!", "love")
	p(result)

	result = strings.ContainsAny("success", "xy")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("five", "")
	p(n)

	p(strings.ToLower("Go Python Java"))
	p(strings.ToUpper("Go Python Java"))

	p("go" == "go")
	p("Go" == "go")

	p(strings.ToLower("Go") == strings.ToLower("go"))

	p(strings.EqualFold("Go", "gO"))

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("strings.Split():%#v\n", s)

	s = strings.Split("Go for Go!", "")
	fmt.Printf("strings.Split():%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	j := strings.Join(s, "-")
	fmt.Printf("%T\n", j)
	p(j)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ")
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("...Hello, Gophers!!!?", ".!?")
	fmt.Printf("%q\n", s2)

}
