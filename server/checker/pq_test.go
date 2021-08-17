package checker

import (
	"container/heap"
	"testing"

	"github.com/rallinator7/akita-poker/server/card"
)

// test HandPriorityQueue creation
func TestNewHandPriorityQueue(t *testing.T) {
	hpq := NewHandPriorityQueue()

	if hpq.Len() != 0 || len(hpq.Queue) != 0 {
		t.Fatal("priority queue must be on length 0 when instantiated")
	}
}

//test push
func TestPriorityQueue_Push(t *testing.T) {
	hpq := NewHandPriorityQueue()

	hand := NewPQHand(HighCard.String(), []card.Card{{Face: card.King, Suit: card.Diamond}}, int(HighCard))

	startLength := hpq.Len()
	heap.Push(hpq, hand)
	endLength := hpq.Len()

	if endLength-startLength != 1 {
		t.Fatal("length of priority queue did not increase after push")
	}

}

// test pop
func TestPriorityQueue_Pop(t *testing.T) {
	hpq := NewHandPriorityQueue()

	h1 := NewPQHand(HighCard.String(), []card.Card{{Face: card.King, Suit: card.Diamond}}, int(HighCard))
	h2 := NewPQHand(RoyalFlush.String(), []card.Card{{Face: card.King, Suit: card.Diamond}}, int(RoyalFlush))
	h3 := NewPQHand(Flush.String(), []card.Card{{Face: card.King, Suit: card.Diamond}}, int(Flush))
	messageList := []*PQHand{h1, h2, h3}

	for _, hand := range messageList {
		heap.Push(hpq, hand)
	}

	pushLength := hpq.Len()
	hand := heap.Pop(hpq)
	popLength := hpq.Len()

	if hand.(*PQHand).Name != RoyalFlush.String() {
		t.Fatalf("Expected type of Royal Flush but got: %v", hand.(*PQHand).Name)
	}

	if pushLength-popLength != 1 {
		t.Fatal("Priority Queue did not shrink")
	}
}

// test len
func TestPriorityQueue_Len(t *testing.T) {
	hpq := NewHandPriorityQueue()
	hand := NewPQHand(HighCard.String(), []card.Card{{Face: card.King, Suit: card.Diamond}}, int(HighCard))

	for i := 0; i <= 10; i++ {
		heap.Push(hpq, hand)

		if hpq.Len() != i+1 {
			t.Fatal("length of queue doesn't match amount of pushes")
		}
	}
}

// test empty
func TestPriorityQueue_Empty(t *testing.T) {
	hpq := NewHandPriorityQueue()
	hand := NewPQHand(HighCard.String(), []card.Card{{Face: card.King, Suit: card.Diamond}}, int(HighCard))

	for i := 0; i <= 10; i++ {
		heap.Push(hpq, hand)
	}

	hpq.Empty()

	if hpq.Len() != 0 {
		t.Fatal("length of queue was not changed to zero")
	}
}
