package card

// a playing card struct
// should always have a suit and a value
type Card struct {
	Face Face
	Suit Suit
}

// takes a suit and a value and returns a card
// does not care about what cards need to exist at the same time
// this should be taken care of by functions calling it
func NewCard(suit Suit, face Face) *Card {
	c := Card{
		Suit: suit,
		Face: face,
	}

	return &c
}
