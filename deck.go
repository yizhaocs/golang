package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	fmt.Println("*************************** func newDeck() deck  ***************************")
	cards := deck{}
	// clubs (♣), diamonds (♦), hearts (♥) and spades (♠)
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	/*
		replace i and j with the underscore which tells go,
		hey, we understand that there is a variable here, but we just don't care
	*/
	// for i, suit := range cardSuits {
	for _, suit := range cardSuits {
		// for j, value := range cardValues {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}
func (d deck) print() {
	fmt.Println("*************************** func (d deck) print() ***************************")
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	fmt.Println("*************************** func deal(d deck, handSize int) (deck, deck) ***************************")
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	fmt.Println("*************************** func (d deck) toString() string ***************************")
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	fmt.Println("*************************** func (d deck) saveToFile(filename string) error ***************************")
	// 0666 means anyone can read and write this file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fmt.Println("*************************** func newDeckFromFile(filename string) deck ***************************")
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
