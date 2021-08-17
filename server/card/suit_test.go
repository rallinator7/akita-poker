package card

import "testing"

func TestSuits(t *testing.T) {
	suits := Suits()

	if len(suits) != 4 {
		t.Fatalf("Expected 4 suits but got %v", len(suits))
	}
}

func TestSuit_String(t *testing.T) {
	testSuits := []string{
		"Hearts",
		"Diamonds",
		"Spades",
		"Clubs",
	}

	suits := Suits()

	for i, suit := range suits {
		if suit.String() != testSuits[i] {
			t.Fatalf("Expected %s but got %s", testSuits[i], suit.String())
		}
	}
}
