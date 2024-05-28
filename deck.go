package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Deck []string

func newDeck() Deck {
	cards := Deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}

	return cards
}

func (d Deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(currentDeck Deck, handSize int) (Deck, Deck) {
	return currentDeck[:handSize], currentDeck[handSize:]
}

func (deck Deck) toString() string {
	deckStringSlice := []string(deck)
	deckString := strings.Join(deckStringSlice, ",")
	return deckString
}

func (deck Deck) saveToFile(name string) error {
	return os.WriteFile(name, []byte(deck.toString()), 0666)
}

func toDeckFromByteSlice(data []byte) Deck {
	deckSlice := strings.Split(string(data), ",")
	return Deck(deckSlice)
}

func readFromFile(name string) Deck {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Failed to load existing deck. Creating new deck")
		return newDeck()
	}
	return toDeckFromByteSlice(data)
}

func (deck Deck) shuffle() Deck {
	for index := range deck {
		newRandomNumber := rand.Intn(len(deck) - 1)
		deck[index], deck[newRandomNumber] = deck[newRandomNumber], deck[index]
	}

	return deck
}
