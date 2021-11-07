package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = bookstore.Catalog{
	"1": {
		ID:     "1",
		Title:  "Nicholas Chucklaeby",
		Author: "Charles Dickens",
		Copies: 1,
	},
	"2": {
		ID:     "2",
		Title:  "Delightfully Uneventful Trip on the Orient Express",
		Author: "Agatha Christie",
		Copies: 2,
	},
}

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		ID:     "1",
		Title:  "Nicholas Chucklaeby",
		Author: "Charles Dickens",
		Copies: 1,
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	want := []bookstore.Book{
		{
			ID: "1",
			Title:  "Hungry 56",
			Author: "Andy Anderson",
			Copies: 1,
		},
		{
			ID: "2",
			Title:  "Delightfully Uneventful Trip on the Orient Express",
			Author: "Agatha Christie",
			Copies: 2,
		},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	want := "Nicholas Chucklaeby, by Charles Dickens"
	got := bookstore.GetBookDetails(catalog, "1")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
