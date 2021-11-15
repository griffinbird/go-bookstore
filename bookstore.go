package bookstore

import "fmt"

type Book struct {
	ID string
	Title      string
	Author     string
	Copies     int
	PriceCents int
	DiscountPercent int
}

type Catalog map[string]Book

func GetAllBooks(catalog Catalog) []Book {
	books := []Book{}
	for _, b := range catalog {
		books = append(books, b)
	}
	return books
}

func GetBookDetails(catalog Catalog, ID string) string {
	b := catalog[ID]
	return fmt.Sprintf("%s, by %s", b.Title, b.Author)
}

func (b Book) NetPrice() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b Book) SalesPrice() int {
	saving := b.PriceCents * 50 / 100
	return b.PriceCents - saving
}

func (b Book) Buy() int {
	return 0
}

func (b *Book) SetTitle(t string) {
	b.Title = t
}