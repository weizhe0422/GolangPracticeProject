package main

import (
	"fmt"
)

func main() {
	cards1 := newDeck()
	cards1.savetoFile("allCards")

	cards := newDeckFromFile("allCards")
	fmt.Println(cards.toString())
	cards.shuffle()
	hand, remainCards := cards.deal(5)
	hand.print()
	fmt.Println("-----")
	remainCards.print()
}
