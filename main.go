package main

import (
	"github.com/floydkretschmar/golang_cards/cards"
)

func main() {
	deck := cards.NewDeck()
	deck.Print()
}
