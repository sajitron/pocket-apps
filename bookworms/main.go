package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Printing bookworms")
	var file string
	flag.StringVar(&file, "file", "testdata/bookworms.json", "The path to the bookworms file")
	flag.Parse()

	bookworms, err := loadBookworms(file)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)
}
