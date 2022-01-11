package main

import "fmt"

func main(){
	cards := newDeck()
	// testing newDeck function
	cards.print()
	fmt.Println(cards.toString())
	// testing saveToFile and newDeckFromFile function
	filename := "my_cards.txt"
	cards.saveToFile(filename)
	cards_new := newDeckFromFile(filename)
	cards_new.print()
	// testing shuffle function
	cards.shuffle()
	cards.print()
}
