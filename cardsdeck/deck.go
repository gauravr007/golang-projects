package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Creates a new deck of all 52 cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			card := fmt.Sprintf("%s of %s", suit, cardValue)
			cards = append(cards, card)
		}
	}

	return cards
}

// Deal the given deck of cards to get a hand of specific size.
// Returns both the hand the remaining deck of cards
func deal(d deck, handSize int) (hand deck, remainingCards deck){
	if len(d) > handSize {
		return d[:handSize], d[handSize:]
	}

	return d, deck{}
}

// Prints the values of the cards present in the deck
func (d deck) print() {
	fmt.Println("*******************************************************")
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("*******************************************************")
}

// Shuffle

// Save deck to file

// Read deck from file
