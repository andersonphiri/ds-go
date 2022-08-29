package heaps

import (
	"container/heap"
	"log"
)
const (
	DefaultQueueSize = 8
	MinimumInitialCount = 0
)
type priorityQueue[TKey any] struct {
	heap []interface{}
	less func(left, right *TKey) bool 
	count int 
}
func newPriorityQueue[T any](initialCount int, less func(left, right *T) bool ) *priorityQueue[T] {
	if initialCount < MinimumInitialCount {
		log.Panicf("initial count should be at least one, you have set %d", initialCount)
	}
	pq := &priorityQueue[T]{less: less, count: 0}
	pq.heap = make([]interface{}, 0, initialCount + DefaultQueueSize)
	return pq
}

func newPriorityQueueFromList[T any](list []T, less func(left, right *T) bool ) *priorityQueue[T] {
	
	pq := &priorityQueue[T]{less: less}
	pq.heap = make([]interface{}, 0, len(list) + DefaultQueueSize)
	for _, item := range list {
		pq.heap = append(pq.heap, item)
	}
	
	pq.count = len(list)
	return pq
}

func (pq *priorityQueue[T]) Less(i,j int) bool {
	left := pq.heap[i].(T)
	right := pq.heap[j].(T)
	return pq.less(&left, &right)
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.heap)
}

func (pq *priorityQueue[T]) Swap(i,j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *priorityQueue[T]) Push(item interface{}) {
	pq.heap = append(pq.heap, item)
	pq.count++
}

func (pq *priorityQueue[T]) Pop() interface{} {
	if pq.count == 0 {
		return nil
	}
	n := len(pq.heap)
	toBeDeleted := pq.heap[n - 1]
	pq.heap = pq.heap[0:n - 1]
	pq.count--
	return toBeDeleted
}

type PriorityQueue[TKey any] struct {
	pq *priorityQueue[TKey]
}

// NewPriorityQueue creates an empty Priority queue with a Count() or Len() of zero
func NewPriorityQueue[T any](initialCount int, less func(left, right *T) bool ) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{pq: newPriorityQueue(initialCount, less)}
	return pq
}

// NewPriorityQueueFromList creates an empty Priority queue with a Count() or Len() of zero
func NewPriorityQueueFromList[T any](items []T, less func(left, right *T) bool ) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{pq: newPriorityQueueFromList(items, less)}
	heap.Init(pq.pq)
	return pq
}


// Insert add a new element to the queue, then maintains the heap invariant
// O(logN)
func (pq *PriorityQueue[T]) Insert(item T) {
	heap.Push(pq.pq, item)
}

// Delete removes the minimum or maximum element in the queue
// whether the element is minimum or maximum is determined by the less function you use
// O(logN) , then maintains the heap invariant
// returns nil if queue is empty or cast from interface{} to type fails
func (pq *PriorityQueue[T]) Delete() *T {
	if len(pq.pq.heap) == 0 {
		return nil
	}
	poped := heap.Pop(pq.pq)
	toBeDeleted ,ok := poped.(T)
	if !ok {
		return nil
	} else {
		return &toBeDeleted
	}
}

// MinOrMax returns minimum if this is a max priority queue or returns minimum element if this is a min priority queue
func (pq *PriorityQueue[T]) MinOrMax() *T {
	if len(pq.pq.heap) == 0 {
		return nil
	}
	result , ok := pq.pq.heap[0].(T)
	if !ok {
		return nil
	} else {
		return &result
	}
}

// Count counts the number of items remaining in the queue
// similar to Len
func (pq *PriorityQueue[T]) Count() int {
	return pq.Len()
}


// Len counts the number of items remaining in the queue
// similar to Count
func (pq *PriorityQueue[T]) Len() int {
	if pq == nil {
		return 0
	}
	return pq.pq.Len()
}

