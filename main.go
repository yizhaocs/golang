package main

import "fmt"

func main(){
	cards := newDeck()
	cards.print()
	fmt.Println(cards.toString())
	filename := "my_cards.txt"
	cards.saveToFile(filename)
	cards_new := newDeckFromFile(filename)
	cards_new.print()
}
