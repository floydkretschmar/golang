package cards

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
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

func NewDeckFromFile(filename string) Deck {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	s := string(data)
	return strings.Split(s, ";")
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func Shuffle(d Deck) Deck {
	shuffledDeck := make(Deck, len(d))
	copy(shuffledDeck, d)

	shuffledAtLeastOne := false
	for i, card := range shuffledDeck {
		pos := rand.Intn(len(shuffledDeck))
		shuffledAtLeastOne = shuffledAtLeastOne || pos != i
		cardAtPos := shuffledDeck[pos]
		shuffledDeck[pos] = card
		shuffledDeck[i] = cardAtPos
	}

	if shuffledAtLeastOne {
		return shuffledDeck
	} else {
		return Shuffle(shuffledDeck)
	}
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) Save(fileName string, permissions fs.FileMode) {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), permissions)
	if err != nil {
		log.Panicf("The file %s cannot be saved: %s", fileName, err.Error())
	}
}

func (d Deck) toString() string {
	return strings.Join(d, ";")
}
