package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Γειά σου κόσμο",    // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "سلام دنیا",         // Urdu
	"vi": "Xin chào thế giới", // Vietnamese
}

// greet says hello to the world in the specified language
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
