package main

func main() {
	cards := newDeck()

	hand, cards := deal(cards, 7)
	hand.print()
	cards.print()
}
