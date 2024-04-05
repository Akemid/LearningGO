package main

import (
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filePath string) error {
	text := d.toString()
	return os.WriteFile(filePath, []byte(text), 0666)
}

func readFromFile(filePath string) deck {
	bs,error := os.ReadFile(filePath)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	deckString := string(bs)
	deckSlice := strings.Split(deckString, ",")
	returnDeck := deck(deckSlice)
	return returnDeck
}

func (d deck) shuffle() {
	seedValue := time.Now().UnixNano()
	source := rand.NewSource(seedValue)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}