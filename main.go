package main

import (
	"deckofcard/cards"
	"fmt"
)

func main() {

	deckvar := cards.Newdeck()
	fmt.Println("**********************Display cards***************************")
	deckvar.Displaycards()
	fmt.Println("**********************Display shuffle cards******************************")
	deckvar.Shuffel().Displaycards()
	hand, remaining := cards.Drawcards(deckvar, 7)
	fmt.Println("****************************Display inhand cards***************************************")
	hand.Displaycards()
	fmt.Println("******************************Display remaining cards***************************")
	remaining.Displaycards()

}
