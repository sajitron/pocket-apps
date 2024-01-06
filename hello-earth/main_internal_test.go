package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"Greek": {
			lang: "el",
			want: "Γειά σου κόσμο",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Hebrew": {
			lang: "he",
			want: "שלום עולם",
		},
		"Urdu": {
			lang: "ur",
			want: "سلام دنیا",
		},
		"Vietnamese": {
			lang: "vi",
			want: "Xin chào thế giới",
		},
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	// range over all the scenarios
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
