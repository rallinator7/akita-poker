package checker

import (
	"testing"

	"github.com/rallinator7/akita-poker/server/card"
)

// mostly for code coverage percentage
// internal structs have their own tests
func TestNewHandChecker(t *testing.T) {
	NewHandChecker()
}

// struct for CheckHand test table
type CheckHandTest struct {
	Cards           []card.Card
	ExpectedOutPut  Hand
	ExpectedFailure bool
}

// tests to make sure a single HandChecker can work over multiple card hands
// checks for correct hand length, card uniqueness, and hand check accuracy
func TestHandChecker_CheckHand(t *testing.T) {
	hc := NewHandChecker()

	handTestTable := []CheckHandTest{
		{
			Cards: []card.Card{
				{Face: card.Ace, Suit: card.Spade},
				{Face: card.Queen, Suit: card.Spade},
				{Face: card.King, Suit: card.Spade},
				{Face: card.Jack, Suit: card.Spade},
				{Face: card.Ten, Suit: card.Spade},
			},
			ExpectedOutPut:  RoyalFlush,
			ExpectedFailure: false,
		},
		{
			Cards: []card.Card{
				{Face: card.Six, Suit: card.Spade},
				{Face: card.Two, Suit: card.Spade},
				{Face: card.Three, Suit: card.Spade},
				{Face: card.Four, Suit: card.Spade},
				{Face: card.Five, Suit: card.Spade},
			},
			ExpectedOutPut:  StraightFlush,
			ExpectedFailure: false,
		},
		{
			Cards: []card.Card{
				{Face: card.Six, Suit: card.Spade},
				{Face: card.Two, Suit: card.Diamond},
				{Face: card.King, Suit: card.Spade},
				{Face: card.Five, Suit: card.Spade},
				{Face: card.Five, Suit: card.Spade},
			},
			ExpectedFailure: true,
		},
		{
			Cards: []card.Card{
				{Face: card.Six, Suit: card.Spade},
				{Face: card.Two, Suit: card.Diamond},
				{Face: card.King, Suit: card.Spade},
				{Face: card.Five, Suit: card.Spade},
			},
			ExpectedFailure: true,
		},
	}

	for _, test := range handTestTable {
		name, _, err := hc.CheckHand(test.Cards)
		if err != nil {
			if test.ExpectedFailure {
				continue
			}
			t.Fatalf("Expected hand check to pass but it failed: %s", err)
		}

		if name != test.ExpectedOutPut.String() {
			t.Fatalf("Expected %s but got %s", test.ExpectedOutPut, name)
		}
	}
}
