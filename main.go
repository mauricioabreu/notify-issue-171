package main

import (
	"log"

	"github.com/rjeczalik/notify"
)

func main() {
	c := make(chan notify.EventInfo, 1)
	err := notify.Watch("example/...", c, notify.All)
	if err != nil {
		panic(err)
	}

	defer notify.Stop(c)

	for {
		ei := <-c
		log.Printf("New event %s %v\n", ei.Path(), ei.Event())
	}
}
