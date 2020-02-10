package main

import (
	"log"
	"sync"

	"github.com/rjeczalik/notify"
)

func main() {
	c, err := QueuedWatch("example/...", notify.All)
	if err != nil {
		log.Fatalln(err)
	}
	for ev := range c {
		log.Printf("New event %s %v\n", ev.Path(), ev.Event())
	}
}

func QueuedWatch(path string, events ...notify.Event) (chan notify.EventInfo, error) {
	in := make(chan notify.EventInfo, 10240)
	err := notify.Watch(path, in, events...)
	if err != nil {
		return nil, err
	}
	out := make(chan notify.EventInfo)
	queueEvents(in, out)
	return out, nil
}

type eventQueue struct {
	mu    sync.Mutex
	queue []notify.EventInfo
	sleep chan struct{}
}

func queueEvents(in, out chan notify.EventInfo) {
	ev := &eventQueue{
		sleep: make(chan struct{}, 1),
	}
	go ev.enqueueLoop(in)
	go ev.popLoop(out)
}

func (eq *eventQueue) enqueueLoop(in chan notify.EventInfo) {
	var sleeping bool
	for ev := range in {
		eq.mu.Lock()
		sleeping = len(eq.queue) == 0
		eq.queue = append(eq.queue, ev)
		if sleeping {
			select {
			case eq.sleep <- struct{}{}:
			default:
			}
		}
		eq.mu.Unlock()
	}
}

func (eq *eventQueue) pop() (ev notify.EventInfo, n int) {
	eq.mu.Lock()
	defer eq.mu.Unlock()
	if n = len(eq.queue); n == 0 {
		return nil, 0
	}
	ev, eq.queue = eq.queue[0], eq.queue[1:]
	return ev, n
}

func (eq *eventQueue) popLoop(out chan notify.EventInfo) {
	var ev notify.EventInfo
	var n int
	for range eq.sleep {
		for ev, n = eq.pop(); n != 0; ev, n = eq.pop() {
			out <- ev
		}
	}
}
