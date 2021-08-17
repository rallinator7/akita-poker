package checker

import (
	"container/heap"
)

type PriorityQueuer interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
	Push(interface{})
	Pop() interface{}
	Empty()
}

// a priority queue that holds messages
type HandPriorityQueue struct {
	Queue []*PQHand
	lock  Locker
}

//creates a new HandPriorityQueue
func NewHandPriorityQueue() *HandPriorityQueue {
	hpq := HandPriorityQueue{
		Queue: make([]*PQHand, 0),
		lock:  NewTicketLock(),
	}

	heap.Init(&hpq)
	return &hpq
}

// gets length of the queue
func (hpq *HandPriorityQueue) Len() int { return len(hpq.Queue) }

// needed to implemented heap interface and is used for comparing nodes in queue
func (hpq *HandPriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return hpq.Queue[i].Priority > hpq.Queue[j].Priority
}

// needed to implemented heap interface and is used for swapping nodes in queue
func (hpq *HandPriorityQueue) Swap(i, j int) {
	hpq.Queue[i], hpq.Queue[j] = hpq.Queue[j], hpq.Queue[i]
	hpq.Queue[i].Index = i
	hpq.Queue[j].Index = j
}

// retrieves a lock and adds a node into the queue
func (hpq *HandPriorityQueue) Push(message interface{}) {
	hpq.lock.Lock()
	defer hpq.lock.Unlock()

	n := hpq.Len()
	hpqHand := message.(*PQHand)
	hpqHand.Index = n
	hpq.Queue = append(hpq.Queue, hpqHand)
}

// retrieves a lock and pops highest priority node from queue
func (hpq *HandPriorityQueue) Pop() interface{} {
	hpq.lock.Lock()
	defer hpq.lock.Unlock()

	old := hpq.Queue
	n := hpq.Len()
	hpqHand := old[n-1]
	old[n-1] = nil     // avoid memory leak
	hpqHand.Index = -1 // for safety
	hpq.Queue = old[0 : n-1]
	return hpqHand
}

// empties the queue
func (hpq *HandPriorityQueue) Empty() {
	hpq.lock.Lock()
	defer hpq.lock.Unlock()
	hpq.Queue = []*PQHand{}
}
