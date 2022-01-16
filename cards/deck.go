package cards

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
)

type Deck []string

func NewDeck() Deck {
	deck := Deck{}
	cardSuits := []string{"Clubs", "Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			deck = append(deck, fmt.Sprintf("%s of %s", cardValue, cardSuit))
		}
	}
	return deck
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) Deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) Save(fileName string, permissions fs.FileMode) {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), permissions)
	if err != nil {
		log.Fatal(err)
	}
}

func (d Deck) toString() string {
	deckString := d[0]
	for _, card := range d[1:] {
		deckString += "\n" + card
	}
	return deckString
}
