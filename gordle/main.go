package main

import (
	"os"

	"github.com/sajitron/pocket-apps/gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
