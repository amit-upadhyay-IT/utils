package pq

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    interface{}
	priority int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	// TODO: find a way to update the less method
	less  func(i, j int) bool
	queue []*Item
}

func (pq PriorityQueue) Len() int { return len(pq.queue) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.queue[i].priority > pq.queue[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.queue[i].index = i
	pq.queue[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(pq.queue)
	item := x.(*Item)
	item.index = n
	pq.queue = append(pq.queue, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old.queue)
	item := old.queue[n-1]
	old.queue[n-1] = nil // avoid memory leak
	item.index = -1      // for safety
	pq.queue = old.queue[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
