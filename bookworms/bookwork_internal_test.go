package main

import (
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	villette      = Book{Author: "Charlotte Brontë", Title: "Villette"}
	ilPrincipe    = Book{Author: "Niccolò Machiavelli", Title: "Il Principe"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantError     bool
	}

	var tests = map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantError: false,
		},
		"file does not exist": {
			bookwormsFile: "testdata/does-not-exist.json",
			want:          nil,
			wantError:     true,
		},
		"invalid file": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantError:     true,
		},
	}

	// range over all the scenarios
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if tc.wantError && err == nil {
				t.Fatal("expected error, got nil")
			}

			if err != nil && !tc.wantError {
				t.Fatalf("expected no error, got %s", err.Error())
			}

			if !equalBookworms(got, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"empty bookworms": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{handmaidsTale: 1, theBellJar: 1},
		},
		"bookworms with twice the same books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, handmaidsTale}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 3,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)

			if !equalBooksCount(t, got, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: []Book{handmaidsTale},
		},
		"three bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
				{Name: "Jane", Books: []Book{handmaidsTale, janeEyre}},
			},
			want: []Book{janeEyre, handmaidsTale},
		},
		"output is sorted by authors and then title": {
			input: []Bookworm{
				{Name: "Peggy", Books: []Book{ilPrincipe, janeEyre, villette}},
				{Name: "Did", Books: []Book{janeEyre}},
				{Name: "Ali", Books: []Book{villette, ilPrincipe}},
			},
			want: []Book{janeEyre, villette, ilPrincipe},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)

			if !equalBooks(got, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(bookworms, target []Bookworm) bool {
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(books, target []Book) bool {
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}

// equalBooksCount is a helper to test the equality of two maps of books count.
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}
	return true
}
