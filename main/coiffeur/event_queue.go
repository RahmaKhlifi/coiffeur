package coiffeur

import (
	"container/heap"
)

// Event represents an event with a time.
type Event struct {
	event    interface{}
	priority float64 // The time  (lower value = higher priority).
	index    int     // The index of the task in the heap.
}

func (this *Event) equals(other *Event) bool {
	return this.event == other.event && this.priority == other.priority
}

type EventQueue struct {
	pq *InternalEventQueue
}

func NewEventQueue() *EventQueue {
	pq := make(InternalEventQueue, 0)
	heap.Init(&pq)
	return &EventQueue{&pq}
}

func (this *EventQueue) Push(e *Event) {
	heap.Push(this.pq, e)

}

func (this *EventQueue) Pop() *Event {
	res := heap.Pop(this.pq).(*Event)
	return res
}

func (this *EventQueue) Len() int {
	return this.pq.Len()
}

// EventQueue implements a priority queue of Events.
type InternalEventQueue []*Event

// Len is part of heap.Interface.
func (pq *InternalEventQueue) Len() int { return len(*pq) }

// Less is part of heap.Interface. Higher-priority tasks are "less".
func (pq *InternalEventQueue) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

// Swap is part of heap.Interface.
func (pq *InternalEventQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

// Push adds an element to the heap.
func (pq *InternalEventQueue) Push(x interface{}) {
	task := x.(*Event)
	task.index = len(*pq)
	*pq = append(*pq, task)
}

// Pop removes and returns the element with the highest priority

func (pq *InternalEventQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil  // Avoid memory leak
	task.index = -1 // For safety
	*pq = old[0 : n-1]
	return task
}

// Update modifies the time of a task in the queue.
func (pq *InternalEventQueue) Update(task *Event, priority float64) {
	task.priority = priority
	heap.Fix(pq, task.index)
}
