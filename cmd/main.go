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

func Overwrite(s []int) {
	s[0] = 0
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(nums)
	Overwrite(nums)
	fmt.Println(nums)
}
