package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Game holds all the information we need to play a game of Gordle.
type Game struct {
	reader *bufio.Reader
}

const solutionLength = 5

// New returns a game which can be used to play.
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}
	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	// ask for a valid word
	guess := g.ask()

	fmt.Printf("Your guess is %s\n", string(guess))
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		// In order for any non-ascii character to be read, we need to convert the input to a string.
		// before converting to a rune slice.
		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution! Expected %d characters, got %d.\n", solutionLength, len(guess))
		} else {
			return guess
		}
	}
}
