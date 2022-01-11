package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	fmt.Println("################# test case 1 begin #################")
	if len(d) != 52 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	fmt.Println("################# test case 2 begin #################")
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	fmt.Println("################# test case 3 begin #################")
	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", d[len(d) - 1])
	}
}
