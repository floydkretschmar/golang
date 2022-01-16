package cards

import (
	"bou.ke/monkey"
	"errors"
	"io/ioutil"
	"log"
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
	checkIfDecksAreEquivalent(t, deck, expectedDeck)
}

func TestDeck_Deal(t *testing.T) {
	deck := NewDeck()
	handSize := 5
	hand, leftoverDeck := Deal(deck, handSize)

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
	expectedDeckString := "Ace of Clubs;2 of Clubs;3 of Clubs;4 of Clubs;5 of Clubs;6 of Clubs;7 of Clubs;8 of Clubs;9 of Clubs;10 of Clubs;Jack of Clubs;Queen of Clubs;King of Clubs;" +
		"Ace of Spades;2 of Spades;3 of Spades;4 of Spades;5 of Spades;6 of Spades;7 of Spades;8 of Spades;9 of Spades;10 of Spades;Jack of Spades;Queen of Spades;King of Spades;" +
		"Ace of Diamonds;2 of Diamonds;3 of Diamonds;4 of Diamonds;5 of Diamonds;6 of Diamonds;7 of Diamonds;8 of Diamonds;9 of Diamonds;10 of Diamonds;Jack of Diamonds;Queen of Diamonds;King of Diamonds;" +
		"Ace of Hearts;2 of Hearts;3 of Hearts;4 of Hearts;5 of Hearts;6 of Hearts;7 of Hearts;8 of Hearts;9 of Hearts;10 of Hearts;Jack of Hearts;Queen of Hearts;King of Hearts"
	deckString := deck.toString()

	if expectedDeckString != deckString {
		t.Errorf("Expected deck string to be \n%s\n but got \n%s\n instead.", expectedDeckString, deckString)
	}
}

func TestDeck_Save(t *testing.T) {
	deck := NewDeck()
	fileName := "deck"

	err := trySave(deck, fileName)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ioutil.ReadFile(fileName)

	if err != nil {
		t.Errorf(err.Error())
	} else {
		deleteDeck(fileName, t)
	}
}

func TestDeck_NewDeckFromFile(t *testing.T) {
	deck := NewDeck()
	fileName := "deck"
	createdDeck := false

	err := trySave(deck, fileName)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		createdDeck = true
	}

	deckFromFile := NewDeckFromFile(fileName)
	checkIfDecksAreEquivalent(t, deckFromFile, deck)

	if createdDeck {
		deleteDeck(fileName, t)
	}
}

func TestDeck_NewDeckFromFileDeckDoesNotExist(t *testing.T) {
	fatalCalled := false
	fakeLogFatal := func(msg ...interface{}) {
		t.Log(msg)
		fatalCalled = true
	}
	patch := monkey.Patch(log.Fatal, fakeLogFatal)
	defer patch.Unpatch()
	NewDeckFromFile("doesnotexist")

	if !fatalCalled {
		t.Errorf("A fatal error should have ocurred while reading the file but did not.")
	}
}

func TestDeck_Shuffle(t *testing.T) {
	deck := NewDeck()
	originalDeck := make(Deck, len(deck))
	copy(originalDeck, deck)
	deck.Shuffle()

	exactSameOrder := true
	for i, shuffledCard := range deck {
		if originalDeck[i] != shuffledCard {
			exactSameOrder = false
			break
		}
	}

	if exactSameOrder {
		t.Errorf("Expected the shuffled deck to be different from the provided deck but the decks had the equivalent order.")
	}
}

func deleteDeck(filename string, t *testing.T) {
	err := os.Remove(filename)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func checkIfDecksAreEquivalent(t *testing.T, deck Deck, expectedDeck Deck) {
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
			t.Errorf("The created deck does not contain the card: %s", notFoundCard)
		}
	}
}

func trySave(d Deck, filename string) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	d.Save(filename, 0774)
	return err
}
