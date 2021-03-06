package main

import "fmt"

// Rank

type Rank int

const MaxRank = Rank(13)

func WalkRanks(f func(rank Rank)) {
	for i := Rank(1); i <= MaxRank; i++ {
		f(Rank(i))
	}
}

// Color

type Color int

const (
	Black = Color(iota)
	Red
)

// Suit

type Suit int

const (
	Spades = Suit(iota)
	Hearts
	Diamonds
	Clubs
	LimitSuit
)
const NumSuits = int(LimitSuit)

func WalkSuits(f func(Suit)) {
	for s := Spades; s <= Clubs; s++ {
		f(s)
	}
}

func (s Suit) String() string {
	switch s {
	case Spades:
		return "\u2660"
	case Hearts:
		return "\u2661"
	case Diamonds:
		return "\u2662"
	case Clubs:
		return "\u2663"
	default:
		panic(fmt.Errorf("Invalid suit value: %d", int(s)))
	}
}

func (s Suit) Color() Color {
	switch s {
	case Spades:
		return Black
	case Hearts:
		return Red
	case Diamonds:
		return Red
	case Clubs:
		return Black
	default:
		panic(fmt.Errorf("Invalid suit value: %d", int(s)))
	}
}

// Card

type Card struct {
	Rank
	Suit
}

func WalkCards(f func(card Card)) {
	WalkRanks(func(rank Rank) {
		WalkSuits(func(suit Suit) {
			f(Card{rank, suit})
		})
	})
}

// Card.String returns a string representation of the Card, showing its rank and
// suit, with padding to the left as required to make it three characters
// wide.
func (c Card) String() string {
	switch c.Rank {
	case 1:
		return fmt.Sprintf(" A%s", c.Suit)
	case 2, 3, 4, 5, 6, 7, 8, 9:
		return fmt.Sprintf(" %d%s", c.Rank, c.Suit)
	case 10:
		return fmt.Sprintf("10%s", c.Suit)
	case 11:
		return fmt.Sprintf(" J%s", c.Suit)
	case 12:
		return fmt.Sprintf(" Q%s", c.Suit)
	case 13:
		return fmt.Sprintf(" K%s", c.Suit)
	default:
		panic(fmt.Errorf("Invalid rank value: %d", int(c.Rank)))
	}
}
