package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			cards = append(cards, value+" of "+suit)
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
	bytes, error := ioutil.ReadFile(filename)
	if error != nil { // 如果出错为nil则print出错误并且退出程序
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	s := strings.Split(string(bytes), ",")
	return deck(s)
}

func (d deck) shuffle() {
	fmt.Println("*************************** func (d deck) shuffle() ***************************")
	/*
		So we're using "seed := time.Now().UnixNano()" to generate a different 64 number.
		因为"func (t Time) UnixNano() int64"的返回值是int64
		Every single time we start up our program, we use that as the seed to generate a new source object,
	*/
	seed := time.Now().UnixNano()
	/*
		and then we use that source object as the basis for our new random number generator.
	*/
	rand.Seed(seed)

	for index := range d {
		newPosition := rand.Intn(len(d) - 1)
		/*
			So we're essentially saying take whatever is that new position and assign it to index
			and then take whatever at index assign it to new position.
		*/
		d[index], d[newPosition] = d[newPosition], d[index] // fancy swap in golang
	}
}

//
//func (d deck) shuffleUdemyTaught() {
//	fmt.Println("*************************** func (d deck) shuffle() ***************************")
//	/*
//		So we're using "seed := time.Now().UnixNano()" to generate a different 64 number.
//		因为"func (t Time) UnixNano() int64"的返回值是int64
//		Every single time we start up our program, we use that as the seed to generate a new source object,
//	 */
//	seed := time.Now().UnixNano()
//	/*
//		and then we use that source object as the basis for our new random number generator.
//	 */
//	source := rand.NewSource(seed)
//	randomNumberGenerator := rand.New(source)
//
//	for index := range d {
//		newPosition := randomNumberGenerator.Intn(len(d) - 1)
//		/*
//			So we're essentially saying take whatever is that new position and assign it to index
//			and then take whatever at index assign it to new position.
//		 */
//		d[index], d[newPosition] = d[newPosition], d[index] // fancy swap in golang
//	}
//}
