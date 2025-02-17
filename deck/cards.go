package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	Ace = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	Spades = iota
	Hearts
	Diamonds
	Clubs
)

type Card struct {
	Rank int
	Suit int
}

var ranks = []string{
	"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King",
}

var suits = []string{
	"Spades", "Hearts", "Diamonds", "Clubs",
}

type Deck []Card

func New(opts ...func([]Card) []Card) Deck {
	var d Deck
	for suit := Spades; suit <= Clubs; suit++ {
		for rank := Ace; rank <= King; rank++ {
			d = append(d, Card{Rank: rank, Suit: suit})
		}
	}
	return d
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func Shuffle(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
    return cards
}

func (d *Deck) Sort() {
	sort.SliceStable(
		*d,
		func(i, j int) bool {
			if (*d)[i].Rank < (*d)[j].Rank {
				return true
			}
			return (*d)[i].Suit < (*d)[j].Suit
		})
}

func (d Deck) Display() {
	for _, v := range d {
		fmt.Println(ranks[v.Rank] + " of " + suits[v.Suit])
	}
}

func Create(n int) func([]Card) []Card {
	return func(c []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, c...)
		}
		return ret
	}
}
