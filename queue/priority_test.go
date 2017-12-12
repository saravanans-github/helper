package queue

import (
	"container/heap"
	"testing"
)

var pq PriorityQueue

func createHeap() {
	// Some items and their priorities.
	items := map[string]uint{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq = make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
}

func TestLen(t *testing.T) {
	createHeap()
	defer pq.Destroy()

	if pq.Len() != 3 {
		t.FailNow()
	}
}

func TestPop(t *testing.T) {
	createHeap()
	defer pq.Destroy()

	if heap.Pop(&pq).(*Item).value != "pear" {
		t.FailNow()
	}
}
