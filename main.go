package main

import (
	"github.com/floydkretschmar/golang_cards/cards"
)

func main() {
	deck := cards.NewDeck()
	hand, leftoverDeck := cards.Deal(deck, 5)
	deck.Print()
	hand.Print()
	leftoverDeck.Print()
}
