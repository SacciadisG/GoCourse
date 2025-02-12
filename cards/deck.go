package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string //Declaring a new type

// For reference, Go isn't an OO langugae, so there's no objects, constructors, or methods*
func newDeck() deck { //The equivalence of a constructor for our new deck type
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

// For reference, func (parameter type) funcName() {} "assigns" this function to the type
// In this game, print() is *almost* like a method of the deck type => we call it with "deck_name.print()"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Note the two returns types of deal (deck, deck) => in Go, you can return multiple values in one return statement
// Here, we're returning two values of type deck, one of size handSize, and one with the remainder of the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// To write to File, we need a "byte slice"; we get a byte slice by converting to slice, and then conv. to byte slice (byte[])
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
