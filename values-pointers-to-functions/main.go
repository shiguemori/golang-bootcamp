package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.5
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {

	*quantity = 3
	*price = 500.5
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"

}

func changeProductByPointer(p *Product) {
	(*p).price = 300
	p.productName = "Bicycle"

}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}

}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["x"] = 30

}

func main() {

	quantity, price, name, sold := 5, 300.2, "Laptop", true
	fmt.Println("BEFORE calling changeValues():", quantity, price, name, sold)

	changeValues(quantity, price, name, sold)
	fmt.Println("AFTER calling changeValues():", quantity, price, name, sold)

	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("AFTER calling changeValuesByPointer():", quantity, price, name, sold)

	present := Product{
		price:       100,
		productName: "Watch",
	}

	changeProduct(present)
	fmt.Println(present)

	changeProductByPointer(&present)
	fmt.Println("AFTER calling changeProductByPointer:", present)

	prices := []int{10, 20, 30}

	changeSlice(prices)
	fmt.Println("prices slice after calling changeSlice():", prices)

	myMap := map[string]int{"a": 1, "b": 2}

	changeMap(myMap)
	fmt.Println("myMap after calling changeMap():", myMap)

}
