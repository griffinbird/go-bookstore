package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = bookstore.Catalog{
	"1": {
		ID:         "1",
		Title:      "Hungry 56",
		Author:     "Andy Anderson",
		Copies:     1,
		PriceCents: 100,
	},
	"2": {
		ID:         "2",
		Title:      "Delightfully Uneventful Trip on the Orient Express",
		Author:     "Agatha Christie",
		Copies:     2,
		PriceCents: 100,
	},
}

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		ID:     "1",
		Title:  "Hungry 56, by Andy Anderson",
		Author: "Andy Anderson",
		Copies: 1,
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	want := []bookstore.Book{
		{
			ID:     "1",
			Title:  "Hungry 56",
			Author: "Andy Anderson",
			Copies: 1,
		},
		{
			ID:     "2",
			Title:  "Delightfully Uneventful Trip on the Orient Express",
			Author: "Agatha Christie",
			Copies: 2,
		},
	}
	got := catalog.GetAllBooks()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	b := bookstore.Book{
		ID:     "1",
		Title:  "Hungry 56",
		Author: "Andy Anderson",
		Copies: 1,
	}
	want := []bookstore.Book{
		b,
	}
	catalog.AddBook(b)
	got := catalog.GetAllBooks()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	want := "Hungry 56, by Andy Anderson"
	got := bookstore.GetBookDetails(catalog, "1")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNetPrice(t *testing.T) {
	b := bookstore.Book{
		Title:           "Good Omens",
		PriceCents:      100,
		DiscountPercent: 25,
	}
	want := 75
	got := b.NetPrice()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSalesPrice(t *testing.T) {
	b := bookstore.Book{
		Title:      "Good Omens",
		PriceCents: 100,
	}
	want := 50
	got := b.SalesPrice()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetTitle(t *testing.T) {
	b := bookstore.Book{
		Title: "Good Omens",
	}
	b.SetTitle("Bad Omens")
	want := "Bad Omens"
	got := b.Title
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetPriceCents(t *testing.T) {
	b := bookstore.Book{
		Title:      "Good Omen",
		PriceCents: 100,
	}
	b.SetPrice(200)
	want := 200
	got := b.PriceCents
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

/*
func TestBuy(t *testing.T) {
	t.Parallel()
 	b := bookstore.Book{
 		Title:  "Spark Joy",
 		Author: "Marie Kondō",
 		Copies: 2,
 	}
 	want := 1
 	result := bookstore.Buy(b)
 	got := result.Copies
 	if want != got {
 		t.Errorf("want %d copies after buying 1 copy from a stock of 2, got %d", want, got)
 	}
} */
