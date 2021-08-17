package checker

import (
	"container/heap"
	"log"
	"testing"

	"github.com/rallinator7/akita-poker/server/card"
)

//test high card accuracy
func TestHighCardCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Club},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go HighCardCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)

	if hand.(*PQHand).Cards[0].Face != card.Ace {
		t.Fatalf("Expected and Ace but got %s", hand.(*PQHand).Cards[0].Face)
	}

}

//test pair accuracy
func TestPairCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Ace, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go PairCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)

	if hand.(*PQHand).Name != Pair.String() {
		t.Fatalf("Expected a Pair but got %s", hand.(*PQHand).Name)
	}

	pq.Empty()

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go PairCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test two pair accuracy
func TestTwoPairCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Ace, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Five, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go TwoPairCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != TwoPair.String() {
		t.Fatalf("Expected Two Pair but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go PairCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test three of a kind accuracy
func TestThreeOfAKindCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Ace, Suit: card.Spade},
		{Face: card.Ace, Suit: card.Heart},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go ThreeOfAKindCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != ThreeOfAKind.String() {
		t.Fatalf("Expected Three of a Kind but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go ThreeOfAKindCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test flush accuracy

func TestFlushCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Ace, Suit: card.Spade},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Spade},
		{Face: card.Four, Suit: card.Spade},
		{Face: card.Five, Suit: card.Spade},
	}

	go FlushCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != Flush.String() {
		t.Fatalf("Expected Flush but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go FlushCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test straight accuracy
func TestStraightCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Six, Suit: card.Spade},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Spade},
		{Face: card.Four, Suit: card.Spade},
		{Face: card.Five, Suit: card.Spade},
	}

	go StraightCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != Straight.String() {
		t.Fatalf("Expected Straight but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go StraightCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test full house accuracy

func TestFullHouseCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Six, Suit: card.Spade},
		{Face: card.Six, Suit: card.Diamond},
		{Face: card.Six, Suit: card.Heart},
		{Face: card.Five, Suit: card.Diamond},
		{Face: card.Five, Suit: card.Spade},
	}

	go FullHouseCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != FullHouse.String() {
		t.Fatalf("Expected Full House but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go FullHouseCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test four of a kind accuracy
func TestFourOfAKindCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Ace, Suit: card.Spade},
		{Face: card.Ace, Suit: card.Heart},
		{Face: card.Ace, Suit: card.Diamond},
		{Face: card.Five, Suit: card.Spade},
	}

	go FourOfAKindCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != FourOfAKind.String() {
		t.Fatalf("Expected Four of a Kind but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go FourOfAKindCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test straight flush accuracy

func TestStraightFlushCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Six, Suit: card.Spade},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Spade},
		{Face: card.Four, Suit: card.Spade},
		{Face: card.Five, Suit: card.Spade},
	}

	go StraightFlushCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != StraightFlush.String() {
		t.Fatalf("Expected Straight Flush but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go StraightFlushCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}

//test royal flush accuracy
func TestRoyalFlushCheck(t *testing.T) {
	pq := NewHandPriorityQueue()
	doneChannel := make(chan bool)

	cards := []card.Card{
		{Face: card.Ten, Suit: card.Spade},
		{Face: card.Jack, Suit: card.Spade},
		{Face: card.Queen, Suit: card.Spade},
		{Face: card.King, Suit: card.Spade},
		{Face: card.Ace, Suit: card.Spade},
	}

	go RoyalFlushCheck(cards, pq, doneChannel)

	<-doneChannel

	hand := heap.Pop(pq)
	pq.Empty()

	if hand.(*PQHand).Name != RoyalFlush.String() {
		t.Fatalf("Expected Royal Flush but got %s", hand.(*PQHand).Name)
	}

	cards1 := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Spade},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}

	go RoyalFlushCheck(cards1, pq, doneChannel)

	<-doneChannel

	if pq.Len() != 0 {
		log.Fatalf("Expected length of pq tobe 0 but got: %v", pq.Len())
	}
}
