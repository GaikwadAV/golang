package cards

import (
	"log"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := Newdeck()

	if len(d) != 52 {
		t.Errorf("expected deck length of 52 but we got %v", len(d)) //check number of cards
	}
	if d[0] != "Two of Heart" {
		t.Errorf("expected result Two of Heart but we got %v", d[0]) //check first card into deck
	}
	if d[len(d)-1] != "Ace of Spade" {
		t.Errorf("expected result Ace of Spade but we got %v", d[len(d)-1]) //check last card into deck
	}
}

func TestDrawcards(t *testing.T) {

	hand, remaining := Drawcards(Newdeck(), 7) //inhand cards and remaining cards
	//hand.Displaycards()
	//remaining.Displaycards()
	if len(hand) != 7 {
		t.Errorf("expected result 15 but we got %v", len(hand)) //check draw cards
	}
	if len(remaining) != 45 {
		t.Errorf("expected result 37 but we got %v", len(remaining)) //remaining card
	}

}

func TestDisplaycards(t *testing.T) {
	display := Newdeck()
	display.Displaycards()
	if len(display) != 52 {
		t.Errorf("expeced result 52 but we got %v", len(display))
	}
}

func TestShuffle(t *testing.T) {
	d := Newdeck()
	if len(d) != 52 {
		t.Errorf("expected result 52 cards and we got %v", len(d))
	}
	newdeckfor := d                    //copy d into newdeckfor
	d.Shuffel()                        //call shuffle method
	checkequal := equal(d, newdeckfor) //call equal method
	log.Print(checkequal)              //print log of checkequal

}

func equal(passeddeck1, passeddeck2 []string) bool { //equal method with parameter(d deck and newdeckforchecking)
	if len(passeddeck1) != len(passeddeck2) { //camparing length
		return false

	}
	for i, value := range passeddeck1 { //iterate using for loop on original deck
		if value != passeddeck2[i] { //comparing 2 deck
			return false
		}
	}
	return true
}
