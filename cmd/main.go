package main

import (
	"fmt"
)

type Book struct {
	Title      string
	Author     string
	Copies     int
	Series     int
	PriceCents int
}

var books []Book

func AddBook(catalog []Book, book Book) []Book {
	catalog = append(catalog, book)
	return catalog
}

func main() {
	books = []Book{
		{
			Title:  "Nicholas Chucklaeby",
			Author: "Charles Dickens",
			Copies: 1,
			Series: 1,
		},
		{
			Title:  "Delightfully Uneventful Trip on the Orient Express",
			Author: "Agatha Christie",
			Copies: 2,
			Series: 1,
		},
	}
	books = AddBook(books, Book{Title: "Spark Joy"})
	for _, b := range books {
		fmt.Println((b.Author))
	}
}
