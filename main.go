package main

import (
	"fmt"
	"github.com/floydkretschmar/golang_cards/maps"
	"github.com/floydkretschmar/golang_cards/website"
)

func main() {
	//evenAndOdd()
	//useMaps()
	website.WebsiteChecker()
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
