package card

import "testing"

// test new card
func TestNewCard(t *testing.T) {
	face := Ace
	suit := Diamond
	c := NewCard(suit, face)

	if c.Face != face {
		t.Fatalf("Expected card face to be %s but got %s", face, c.Face)
	}

	if c.Suit != suit {
		t.Fatalf("Expected card suit to be %s but got %s", suit, c.Suit)
	}
}
