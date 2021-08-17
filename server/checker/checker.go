package checker

import (
	"container/heap"

	"github.com/rallinator7/akita-poker/server/card"
)

type Checker interface {
	CheckHand([]card.Card) (string, []card.Card, error)
}

// used to check poker hands
// finds if a hand of cards is a specific poker hand concurrently
// if a hand is multiple types of hands, it uses a priority queue to find the greatest value
type HandChecker struct {
	HandQueue PriorityQueuer
}

//returns a new HandChecker with a started HandPusher
func NewHandChecker() *HandChecker {
	hpq := NewHandPriorityQueue()
	hc := HandChecker{
		HandQueue: hpq,
	}

	return &hc
}

// should be returned if hand size isn't equal to five
type IncorrectHandSizeError struct{}

func (i *IncorrectHandSizeError) Error() string {
	return "hand must have exactly five cards"
}

// should be returned if there duplicates of
type DuplicateCardError struct{}

func (i *DuplicateCardError) Error() string {
	return "each card in the hand must be unique"
}

// takes a list of cards and returns the name of the best poker hand
func (hc *HandChecker) CheckHand(cards []card.Card) (string, []card.Card, error) {
	doneChannel := make(chan bool, 10)

	// poker hand must be of length 5
	if len(cards) != 5 {
		return "", nil, &IncorrectHandSizeError{}
	}

	valid := verifyUniqueCards(cards)

	// poker hand must have all unique cards
	if !valid {
		return "", nil, &DuplicateCardError{}
	}

	// concurrently check if cards make up each hand
	// pass priority queue to have valid hands pushed to
	// pass doneChannel to alert CheckHand the goroutine has completed
	go HighCardCheck(cards, hc.HandQueue, doneChannel)
	go PairCheck(cards, hc.HandQueue, doneChannel)
	go TwoPairCheck(cards, hc.HandQueue, doneChannel)
	go ThreeOfAKindCheck(cards, hc.HandQueue, doneChannel)
	go FlushCheck(cards, hc.HandQueue, doneChannel)
	go StraightCheck(cards, hc.HandQueue, doneChannel)
	go FullHouseCheck(cards, hc.HandQueue, doneChannel)
	go FourOfAKindCheck(cards, hc.HandQueue, doneChannel)
	go StraightFlushCheck(cards, hc.HandQueue, doneChannel)
	go RoyalFlushCheck(cards, hc.HandQueue, doneChannel)

	done := 0
	// listen on doneChannel until all goroutines send back a done response
	for {
		// probably could remove hard code but not sure if poker will ever have more than 10 hands...
		if done == 10 {
			break
		}
		<-doneChannel
		done++
	}
	close(doneChannel)

	// get the highest value from the priority queue
	hand := heap.Pop(hc.HandQueue)

	// empty queue to avoid spier between runs
	hc.HandQueue.Empty()

	return hand.(*PQHand).Name, hand.(*PQHand).Cards, nil
}

// takes a list of cards and concurrently compares each card to the rest of the cards in the hand
// if all are unique, then it returns true, else it returns false
func verifyUniqueCards(cards []card.Card) bool {
	verifyChannel := make(chan bool)
	defer close(verifyChannel)

	// actually does the uniqueness check
	checkUniqueness := func(card card.Card, hand []card.Card, v chan bool) {
		match := 0
		for _, c := range cards {
			if c.Face == card.Face && c.Suit == card.Suit {
				match++
			}
		}

		// cards still holds the original card because chopping up a slice takes longer than an extra iteration of for loop
		// if two cards are the same then match would be equal to 2
		if match > 1 {
			v <- false
			// return so you don't pass to closed channels!
			return
		}

		v <- true
	}

	// concurrently compares each card to the rest of
	for _, card := range cards {
		go checkUniqueness(card, cards, verifyChannel)
	}

	done := 0
	valid := true

	// while the channel is open, receive bools from goroutines
	for unique := range verifyChannel {
		if !unique {
			valid = false
		}
		done++
		// break once all 5 cards are checked
		if done == len(cards) {
			break
		}
	}

	return valid
}
