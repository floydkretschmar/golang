package cards

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDeck_NewDeck(t *testing.T) {
	expectedDeck := []string{
		"Ace of Spades", "2 of Spades", "3 of Spades", "4 of Spades", "5 of Spades", "6 of Spades", "7 of Spades", "8 of Spades", "9 of Spades", "10 of Spades", "Jack of Spades", "Queen of Spades", "King of Spades",
		"Ace of Hearts", "2 of Hearts", "3 of Hearts", "4 of Hearts", "5 of Hearts", "6 of Hearts", "7 of Hearts", "8 of Hearts", "9 of Hearts", "10 of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts",
		"Ace of Clubs", "2 of Clubs", "3 of Clubs", "4 of Clubs", "5 of Clubs", "6 of Clubs", "7 of Clubs", "8 of Clubs", "9 of Clubs", "10 of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs",
		"Ace of Diamonds", "2 of Diamonds", "3 of Diamonds", "4 of Diamonds", "5 of Diamonds", "6 of Diamonds", "7 of Diamonds", "8 of Diamonds", "9 of Diamonds", "10 of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds",
	}
	deck := NewDeck()

	if len(deck) != len(expectedDeck) {
		t.Errorf("The created deck has length of %d instead of the expected lenght of %d", len(deck), len(expectedDeck))
	}

	var notFoundCards Deck
	for _, expectedCard := range expectedDeck {
		found := false
		for _, card := range deck {
			if card == expectedCard {
				found = true
				continue
			}
		}
		if !found {
			notFoundCards = append(notFoundCards, expectedCard)
		}
	}

	if len(notFoundCards) > 0 {
		for _, notFoundCard := range notFoundCards {
			t.Errorf("The created deck does notFoundCard contain the card: %s", notFoundCard)
		}
	}
}

func TestDeck_Deal(t *testing.T) {
	deck := NewDeck()
	handSize := 5
	hand, leftoverDeck := deck.Deal(handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand to have size %d but received %d cards.", handSize, len(hand))
	}

	for _, handCard := range hand {
		for _, deckCard := range leftoverDeck {
			if handCard == deckCard {
				t.Errorf("Hand card %s still exists in leftover deck.", handCard)
			}
		}
	}
}

func TestDeck_ToString(t *testing.T) {
	deck := NewDeck()
	expectedDeckString := "Ace of Clubs\n2 of Clubs\n3 of Clubs\n4 of Clubs\n5 of Clubs\n6 of Clubs\n7 of Clubs\n8 of Clubs\n9 of Clubs\n10 of Clubs\nJack of Clubs\nQueen of Clubs\nKing of Clubs\n" +
		"Ace of Spades\n2 of Spades\n3 of Spades\n4 of Spades\n5 of Spades\n6 of Spades\n7 of Spades\n8 of Spades\n9 of Spades\n10 of Spades\nJack of Spades\nQueen of Spades\nKing of Spades\n" +
		"Ace of Diamonds\n2 of Diamonds\n3 of Diamonds\n4 of Diamonds\n5 of Diamonds\n6 of Diamonds\n7 of Diamonds\n8 of Diamonds\n9 of Diamonds\n10 of Diamonds\nJack of Diamonds\nQueen of Diamonds\nKing of Diamonds\n" +
		"Ace of Hearts\n2 of Hearts\n3 of Hearts\n4 of Hearts\n5 of Hearts\n6 of Hearts\n7 of Hearts\n8 of Hearts\n9 of Hearts\n10 of Hearts\nJack of Hearts\nQueen of Hearts\nKing of Hearts"
	deckString := deck.toString()

	if expectedDeckString != deckString {
		t.Errorf("Expected deck string to be \n%s\n but got \n%s\n instead.", expectedDeckString, deckString)
	}
}

func TestDeck_Save(t *testing.T) {
	deck := NewDeck()
	fileName := "deck"
	deck.Save(fileName, 0774)

	_, err := ioutil.ReadFile(fileName)

	if err != nil {
		t.Errorf(err.Error())
	} else {
		err = os.Remove(fileName)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}
