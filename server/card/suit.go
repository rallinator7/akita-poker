package card

// the suits a card can be
type Suit int

const (
	Heart Suit = iota
	Diamond
	Spade
	Club
)

func Suits() []Suit {
	return []Suit{
		Heart,
		Diamond,
		Spade,
		Club,
	}
}

// needed to implement Stringer interface
// allows for the string value of the suit to be printed out instead of the int value
func (suit Suit) String() string {
	suits := []string{
		"Hearts",
		"Diamonds",
		"Spades",
		"Clubs",
	}

	return suits[int(suit)]
}
