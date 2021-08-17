package checker

import (
	"container/heap"
	"math/rand"
	"time"

	"github.com/rallinator7/akita-poker/server/card"
)

type Hand int

const (
	HighCard Hand = iota
	Pair
	TwoPair
	ThreeOfAKind
	Flush
	Straight
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// allows for the string value of the suit to be printed out instead of the int value
func (hand Hand) String() string {
	faces := []string{
		"High Card",
		"Pair",
		"Two Pair",
		"Three Of A Kind",
		"Flush",
		"Straight",
		"Full House",
		"Four Of A Kind",
		"Straight Flush",
		"Royal Flush",
	}

	return faces[int(hand)]
}

// finds the largest card in the hand and pushes it through the provided channel
func HighCardCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	high := cards[0]

	// gets the card with the highest value
	for _, card := range cards {
		if card.Face >= high.Face {
			high = card
		}
	}

	hand := NewPQHand(HighCard.String(), []card.Card{high}, int(HighCard))

	// always pushes a hand to the priority queue because a poker hand always has a high card
	heap.Push(pq, hand)
	doneChannel <- true
}

// checks to see if two of the cards in the hand have the same face
// always returns the first pair
// if there are two pairs, the two pair function will catch it and get popped from the pq first
func PairCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	faces := card.Faces()

	// initializes a map with a key value pair for each type of card face
	matchingFaces := map[card.Face][]card.Card{}
	for _, face := range faces {
		matchingFaces[face] = []card.Card{}
	}

	// loops through cards and puts each card into the map based on the face it has
	for _, card := range cards {
		matchingFaces[card.Face] = append(matchingFaces[card.Face], card)
	}

	//checks if any of the map values are equal to two (equivalent to a pair)
	// if there's a match, push it to queue
	for _, v := range matchingFaces {
		if len(v) == 2 {
			hand := NewPQHand(Pair.String(), v, int(Pair))
			heap.Push(pq, hand)
		}
	}

	doneChannel <- true
}

// finds if there are two pairs of the same face in the hand
func TwoPairCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	faces := card.Faces()
	pairs := [][]card.Card{}

	// initializes a map with a key value pair for each type of card face
	matchingFaces := map[card.Face][]card.Card{}
	for _, face := range faces {
		matchingFaces[face] = []card.Card{}
	}

	// loops through cards and puts each card into the map based on the face it has
	for _, card := range cards {
		matchingFaces[card.Face] = append(matchingFaces[card.Face], card)
	}

	// checks if any of the map values are equal to two (equivalent to a pair)
	for _, v := range matchingFaces {
		if len(v) == 2 {
			pairs = append(pairs, v)
		}
	}

	// if there are two pairs, push to queue
	if len(pairs) == 2 {
		hand := NewPQHand(TwoPair.String(), append(pairs[0], pairs[1]...), int(TwoPair))
		heap.Push(pq, hand)
	}

	doneChannel <- true
}

// finds if there are three of a kind of a face of a card
func ThreeOfAKindCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	faces := card.Faces()

	// initializes a map with a key value pair for each type of card face
	matchingFaces := map[card.Face][]card.Card{}
	for _, face := range faces {
		matchingFaces[face] = []card.Card{}
	}

	// checks if any of the map values are equal to two (equivalent to a pair)
	for _, card := range cards {
		matchingFaces[card.Face] = append(matchingFaces[card.Face], card)
	}

	//checks if any of the map values are equal to 3 (equivalent to 3oak)
	// if there's a match, push it to queue
	for _, v := range matchingFaces {
		if len(v) == 3 {
			hand := NewPQHand(ThreeOfAKind.String(), v, int(ThreeOfAKind))
			heap.Push(pq, hand)
		}
	}

	doneChannel <- true
}

// finds if the hand is a flush
func FlushCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	suit := cards[0].Suit
	flush := true

	// loops through cards and compares it suit to the first card in the hand
	// if each card has the same suit then flush will remain true
	for _, card := range cards {
		if card.Suit == suit {
			continue
		} else {
			flush = false
			break
		}
	}

	// checks if flush is true and pushes to queue
	if flush {
		hand := NewPQHand(Flush.String(), cards, int(Flush))
		heap.Push(pq, hand)
	}

	doneChannel <- true
}

// finds if the hand is a straight
func StraightCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	// quick sort's worst case is if hand is already in order so we do a shuffle first
	cards = shuffle(cards)
	// does a quick sort on the face of the cards
	cards = quickSortCardFaces(cards)
	straight := true

	// compares the value of the face for each card
	// since a card face is an iota, a straight's cards value would go up by one when in their sorted order
	for i := 0; i < len(cards)-1; i++ {
		if (cards[i+1].Face - cards[i].Face) == 1 {
			continue
		} else {
			straight = false
			break
		}
	}

	// adds to queue if the hand is a straight
	if straight {
		hand := NewPQHand(Straight.String(), cards, int(Straight))
		heap.Push(pq, hand)

	}

	doneChannel <- true
}

// checks if list of cards is a full house
func FullHouseCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	faces := card.Faces()
	tok := false
	pair := false

	// initializes a map with a key value pair for each type of card face
	matchingFaces := map[card.Face][]card.Card{}
	for _, face := range faces {
		matchingFaces[face] = []card.Card{}
	}

	// loops through cards and puts each card into the map based on the face it has
	for _, card := range cards {
		matchingFaces[card.Face] = append(matchingFaces[card.Face], card)
	}

	// checks if a face has two or three cards and then marks pair or tok true
	for _, v := range matchingFaces {
		if len(v) == 3 {
			tok = true
			// don't need to compare two we know it'll be false
			continue
		}
		if len(v) == 2 {
			pair = true
		}
	}

	// if there are a pair and a 3oak push to queue
	if tok && pair {
		hand := NewPQHand(FullHouse.String(), cards, int(FullHouse))
		heap.Push(pq, hand)
	}

	doneChannel <- true
}

// checks if a list of cards is a four of a kind
func FourOfAKindCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	faces := card.Faces()

	// initializes a map with a key value pair for each type of card face
	matchingFaces := map[card.Face][]card.Card{}
	for _, face := range faces {
		matchingFaces[face] = []card.Card{}
	}

	// loops through cards and puts each card into the map based on the face it has
	for _, card := range cards {
		matchingFaces[card.Face] = append(matchingFaces[card.Face], card)
	}

	//if a face has four cards then pushes to queue
	for _, v := range matchingFaces {
		if len(v) == 4 {
			hand := NewPQHand(FourOfAKind.String(), v, int(FourOfAKind))
			heap.Push(pq, hand)
		}
	}

	doneChannel <- true
}

// finds if a list of cards is a straight flush
func StraightFlushCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	// shuffle cards to avoid quick sort worst case
	cards = shuffle(cards)
	// run quicksort on card faces
	cards = quickSortCardFaces(cards)
	suit := cards[0].Suit
	straight := true
	flush := true

	// loops through cards and compares suit to check for flush and compares face values to check for straight
	for i := 0; i < len(cards)-1; i++ {
		if (cards[i+1].Face - cards[i].Face) == 1 {
			if cards[i].Suit == suit {
				continue
			} else {
				flush = false
			}
		} else {
			straight = false
			break
		}
	}

	// if a straight and a flush push to queue
	if straight && flush {
		hand := NewPQHand(StraightFlush.String(), cards, int(StraightFlush))
		heap.Push(pq, hand)
	}

	doneChannel <- true
}

// checks a list of cards for a royal flush
func RoyalFlushCheck(cards []card.Card, pq PriorityQueuer, doneChannel chan bool) {
	cards = shuffle(cards)
	cards = quickSortCardFaces(cards)
	suit := cards[0].Suit
	straight := true
	flush := true

	if cards[0].Face == card.Ten {
		for i := 0; i < len(cards)-1; i++ {
			if (cards[i+1].Face - cards[i].Face) == 1 {
				if cards[i].Suit == suit {
					continue
				} else {
					flush = false
				}
			} else {
				straight = false
				break
			}
		}

		if straight && flush {
			hand := NewPQHand(RoyalFlush.String(), cards, int(RoyalFlush))
			heap.Push(pq, hand)
		}
	}

	doneChannel <- true
}

// implements the quick sort sorting method
func quickSortCardFaces(cards []card.Card) []card.Card {
	// in case list of cards is 1
	if len(cards) < 2 {
		return cards
	}

	left, right := 0, len(cards)-1

	// pick a random element as pivot
	pivot := rand.Int() % len(cards)

	// swap pivot with last element
	cards[pivot], cards[right] = cards[right], cards[pivot]

	// quicksort
	for i := range cards {
		if cards[i].Face < cards[right].Face {
			cards[left], cards[i] = cards[i], cards[left]
			left++
		}
	}

	cards[left], cards[right] = cards[right], cards[left]

	// recursive call for left and right of pivot
	quickSortCardFaces(cards[:left])
	quickSortCardFaces(cards[left+1:])

	return cards
}

// shuffles a list of cards for randomness
func shuffle(cards []card.Card) []card.Card {
	rand.Seed(time.Now().UnixNano()) //needed for randomness in shuffle order
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}
