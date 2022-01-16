package cards

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
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

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) Save(fileName string, permissions fs.FileMode) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), permissions)
}

func (d Deck) toString() string {
	return strings.Join(d, ";")
}
