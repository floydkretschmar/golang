package main

import (
	"fmt"
	"github.com/floydkretschmar/golang_cards/cards"
	"github.com/floydkretschmar/golang_cards/languagebot"
	"github.com/floydkretschmar/golang_cards/maps"
	"github.com/floydkretschmar/golang_cards/structs"
	"log"
)

func main() {
	//playingCards()
	//evenAndOdd()
	//createPerson()
	//useMaps()
	useBots()
}

func evenAndOdd() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%d is even \n", number)
		} else {
			fmt.Printf("%d is odd \n", number)
		}
	}
}

func playingCards() {
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

func createPerson() {
	alex := structs.Person{
		FirstName: "Alex",
		LastName:  "Anderson",
		ContactInfo: structs.ContactInfo{
			Email: "alex.anderson@mai.com",
			Zip:   "12345",
		},
	}

	//Pointer referencing (explicit)
	//refAlex := &alex
	//refAlex.SetName("Jimmy")
	// Pointer referencing (implicit with receiver)
	alex.SetName("Jimmy")

	var alexandra structs.Person
	// default: each field set to default value of string ("")
	alexandra.FirstName = "Alexandra"
	alexandra.LastName = "Anderson"
	alexandra.ContactInfo.Email = "alexandra.anderson@mail.de"

	alex.Print()
	alexandra.Print()
}

func useMaps() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	var colorsEmpty map[string]string
	colorsMake := make(map[string]string)
	colorsMake["white"] = "#ffffff"
	colorsMake["black"] = "#000000"

	delete(colorsMake, "black")

	fmt.Println(colors)
	fmt.Println(colorsEmpty)
	fmt.Println(colorsMake)

	maps.PrintMap(colors)
}

func useBots() {
	var eBot languagebot.EnglishBot
	var sBot languagebot.SpanishBot

	languagebot.PrintGreeting(eBot)
	languagebot.PrintGreeting(sBot)
}
