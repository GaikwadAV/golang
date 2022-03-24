package cards

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	"strconv"
)

type Deck []string

func Newdeck() Deck {

	cards := Deck{}

	number := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	for _, card := range number {
		for _, n := range suits {
			cards = append(cards, card+" of "+n)
		}
	}
	return cards
}

func (d Deck) Displaycards() {
	// for _, j := range d {
	// 	fmt.Println(j)
	// }
	fmt.Println(d)
}

func Drawcards(d Deck, draw int) (Deck, Deck) {
	return d[:draw], d[draw:]
}

func (d Deck) Shuffel() Deck {

	mrand.Seed(generateSeed(d))

	for i := 0; i < len(d); i++ {
		randomnumber := mrand.Intn(i + 1)
		if i != randomnumber {
			d[randomnumber], d[i] = d[i], d[randomnumber]
		}
	}
	return d
}
func generateSeed(d Deck) int64 {
	value, err := crand.Prime(crand.Reader, len(d))
	if err != nil {
		fmt.Print(err)

	}
	x := value.String()
	val, _ := strconv.ParseInt(x, 10, 64)
	return val
}
