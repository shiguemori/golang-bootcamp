package main

import "fmt"

func main() {

	var employees map[string]string

	fmt.Printf("%#v\n", employees)

	fmt.Printf("No. of elements: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"])

	people := map[string]float64{}

	people["John"] = 30.5
	people["Marry"] = 22

	fmt.Printf("%v\n", people)

	map1 := make(map[int]int)
	fmt.Printf("map1: %#v\n", map1)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,

		"CHF": 600,
	}
	fmt.Println(balances)

	m := map[int]int{1: 3, 4: 5, 6: 8}
	_ = m

	balances["USD"] = 500.5
	balances["GBP"] = 800.8
	fmt.Println(balances)

	v, ok := balances["RON"]

	if ok {
		fmt.Println("The RON Balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	fmt.Printf("balances: %v\n", balances)

	delete(balances, "USD")

	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "X"}

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	friends := map[string]int{"Dan": 40, "Maria": 35}

	neighbors := friends

	friends["Dan"] = 30

	fmt.Println(neighbors)

	colleagues := make(map[string]int)

	for k, v := range friends {
		colleagues[k] = v
	}

}
