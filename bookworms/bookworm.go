package main

import (
	"encoding/json"
	"os"
)

// A Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms, and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// Initialize the type in which the file will be decoded.
	var bookworms []Bookworm

	// Decode the file and store the content in the value bookworms.
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

func findCommonBooks(bookworms []Bookworm) []Book {
	// write books for each bookworm into a map to avoid nested loops
	books := make(map[string]int)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			books[book.Title]++
		}
	}
	// find books that are on more than one bookworm's shelf
	var commonBooks []Book
	for title, count := range books {
		if count > 1 {
			commonBooks = append(commonBooks, Book{Title: title})
		}
	}
	return commonBooks
}
