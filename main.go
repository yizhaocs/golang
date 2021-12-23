package main

import "fmt"

func main(){
	cards := newDeck()
	fmt.Println("####################################")
	cards.print()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("####################################")
	hand.print()
	fmt.Println("####################################")
	remainingCards.print()
}
