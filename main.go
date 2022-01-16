package main

import (
	"github.com/floydkretschmar/golang_cards/cards"
)

func main() {
	deck := cards.NewDeck()
	hand, leftoverDeck := deck.Deal(5)
	deck.Print()
	hand.Print()
	leftoverDeck.Print()
}
