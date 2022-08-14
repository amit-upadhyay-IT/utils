package pq

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue_Push(t *testing.T) {

	pq := PriorityQueue{func(i, j int) bool { return i < j }, nil}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{value: "name", priority: 0}
	heap.Push(&pq, item)
	item = &Item{value: 1.2, priority: 1}
	heap.Push(&pq, item)
	item = &Item{value: 1, priority: 3}
	heap.Push(&pq, item)

	// Take the items out; they arrive in decreasing priority order.
	expectedRes := []interface{}{
		1, 1.2, "name",
	}
	i := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		assert.Equal(t, expectedRes[i], item.value)
		i++
	}
}
