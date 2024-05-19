package gordle

import "fmt"

// Game holds all the information we need to play a game of Gordle.
type Game struct{}

// New returns a game which can be used to play.
func New() *Game {
	g := &Game{}
	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Make a guess:\n")
}
