package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func Deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func NewDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func (d deck) NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {
	eiei := rand.NewSource(time.Now().UnixNano())
	lnwza := rand.New(eiei)
	fmt.Println(lnwza)
	for _, haha := range d {
		d[0], d[lnwza.Intn(len(d)-1)] = d[lnwza.Intn(len(d)-1)], haha
	}
	return d
}
