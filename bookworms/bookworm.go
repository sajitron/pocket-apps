package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

// func findCommonBooks(bookworms []Bookworm) []Book {
// 	// write books for each bookworm into a map to avoid nested loops
// 	books := make(map[string]int)
// 	for _, bookworm := range bookworms {
// 		for _, book := range bookworm.Books {
// 			books[book.Title]++
// 		}
// 	}
// 	// find books that are on more than one bookworm's shelf
// 	var commonBooks []Book
// 	for title, count := range books {
// 		if count > 1 {
// 			commonBooks = append(commonBooks, Book{Title: title})
// 		}
// 	}
// 	return commonBooks
// }

// findCommonBooks finds the books that are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	// Register all books on shelves
	booksOnShelves := booksCount(bookworms)

	var commonBooks []Book

	// Find books that are on more than one bookworm's shelf
	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return sortBooks(commonBooks)
}

// booksCount registers all the books and their occurrences on the bookworm's shelf.
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}

// sortBooks sorts the books by Author and then title
func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})
	return books
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
