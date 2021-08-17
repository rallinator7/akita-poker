package checker

import "github.com/rallinator7/akita-poker/server/card"

// The struct that is held inside of the DealerQueue
// name should be the type of hand and cards should be the cards that make up the hand
type PQHand struct {
	Name     string
	Cards    []card.Card
	Priority int // the priority of the hand (better hand == greater priority)
	Index    int // saved for pq
}

// takes in a name, list of cards, and priority and returns a new PQHand
func NewPQHand(name string, cards []card.Card, priority int) *PQHand {
	pqh := PQHand{
		Name:     name,
		Cards:    cards,
		Priority: priority,
	}

	return &pqh
}
