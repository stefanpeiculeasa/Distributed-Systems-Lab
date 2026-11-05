package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() { l.counter++ }
func (l *LamportClock) Send() { l.Tick() }
func (l *LamportClock) Receive(ts int) {
	l.counter = max(l.counter, ts)
	l.Tick()
}

type Message struct{ sender, ts int }
type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// retuneaza true daca a precede b
func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func main() {
	ch := []chan Message{make(chan Message, 1), make(chan Message, 1), make(chan Message, 1)}
	evCh := make(chan Event, 9)

	for i := 0; i < 3; i++ {
		go func(i int) {
			c := &LamportClock{}
			c.counter += rand.Intn(5)
			c.Tick()
			evCh <- Event{i + 1, c.counter, "internal"}
			c.Send()
			ch[(i+1)%3] <- Message{i + 1, c.counter}
			evCh <- Event{i + 1, c.counter, "send"}
			m := <-ch[i]
			c.Receive(m.ts)
			evCh <- Event{i + 1, c.counter, "recv"}
		}(i)
	}
	var events []Event
	for range 9 {
		events = append(events, <-evCh)
	}

	sort.Slice(events, func(a, b int) bool { return less(events[a], events[b]) })
	for _, e := range events {
		fmt.Printf("P%d %-8s %d\n", e.process, e.kind, e.internal_counter)
	}
}
