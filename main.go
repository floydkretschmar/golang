package main

import (
	"github.com/floydkretschmar/golang_cards/cards"
	"log"
)

func main() {
	deck := cards.NewDeck()
	log.Println("New Deck:")
	deck.Print()
	deck.Shuffle()
	log.Println("Shuffled Deck:")
	deck.Print()
	hand, leftoverDeck := cards.Deal(deck, 5)
	log.Println("Hand:")
	hand.Print()
	log.Println("Remaining Deck:")
	leftoverDeck.Print()
}
