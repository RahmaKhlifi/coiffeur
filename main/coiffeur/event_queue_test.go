package coiffeur

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	//"fmt"
)

func TestEventQueue1(t *testing.T) {
	log.Printf("Start TestEventQueue1")
	pq := NewEventQueue()
	event1 := &Event{event: "ev1", priority: 3}
	event2 := &Event{event: "ev2", priority: 1}
	event3 := &Event{event: "ev3", priority: 0}
	pq.Push(event1)
	pq.Push(event2)
	pq.Push(event3)

	expected := make([]*Event, 3)
	expected[0] = event3
	expected[1] = event2
	expected[2] = event1

	got := make([]*Event, 3)

	i := 0
	for pq.Len() > 0 {
		got[i] = pq.Pop()
		i = i + 1
	}

	for _, ev := range expected {
		fmt.Printf("Event: %s, Time: %f\n", ev.event, ev.priority)
	}
	assertEqual(got, expected)
}

// test event queue for 1000 randomly generated events
func TestEventQueue2(t *testing.T) {
	log.Printf("Start TestEventQueue2")

	pq := NewEventQueue()

	i := 0
	for i < 1000 {
		pq.Push(&Event{priority: rand.Float64()})
		i = i + 1
	}
	got := make([]*Event, 0)

	i = 0
	for pq.Len() > 0 {
		got = append(got, pq.Pop())
		i = i + 1

	}

	if i != 1000 {
		log.Fatalf("unexpected length")
	}
	for k := 0; k < 1000-1; k++ {
		if got[k].priority > got[k+1].priority {
			log.Fatalf("Uncorrect order ! ")
		}
	}

}

func assertEqual(expected, actual []*Event) {
	if len(expected) != len(actual) {
		log.Fatalf("length of expected : %v != length of actual : %v\n", len(expected), len(actual))
	}

	for i, _ := range expected {
		if !expected[i].equals(actual[i]) {
			log.Fatalf("Events do not match: %v != %v\n",
				expected[i], actual[i])
		}
	}
}
