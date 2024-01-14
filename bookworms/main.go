package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Printing bookworms")
	bookworms, err := loadBookworms("testdata/bookworms.json")

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)
}