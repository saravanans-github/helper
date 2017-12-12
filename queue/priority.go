package queue

// An Item is something we manage in a priority queue.
//
type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority uint        // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// PriorityQueue is an array of Items managed by the Heap interface
type PriorityQueue []*Item

// Len returns the length of the queue
func (pq PriorityQueue) Len() int { return len(pq) }

// Less is the implementation for the heap.interface to get the highest priority item in the queue
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

// Push is the implementation for heap.interface to add a new item into the Heap
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop is the implementation for heap.interface to get the item with the highest priority
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = 0 // for safety
	*pq = old[0 : n-1]
	return item
}

// Swap is the implementation for heap.interface to swap two items in the heap
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Destroy is the implementation to nil-lify the priority queue for the garbage collector to clear
func (pq PriorityQueue) Destroy() {
	pq = nil
}
