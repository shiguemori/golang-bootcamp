package main

import "fmt"

func main() {

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year, pages   int
	}

	lastBook := book{"The Divine Comedy", "Dante Aligheri", 1320}
	fmt.Println(lastBook)

	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	_ = bestBook

	aBook := book{title: "Just a random book"}
	fmt.Printf("%#v\n", aBook)

	fmt.Println(lastBook.title)

	lastBook.author = "The Best"
	lastBook.year = 2020
	fmt.Printf("lastBook: %+v\n", lastBook)

	randomBook := book{title: "Random Title", author: "John Doe", year: 100}
	fmt.Println(randomBook == lastBook)

	myBook := randomBook
	myBook.year = 2020
	fmt.Println(myBook, randomBook)

}
