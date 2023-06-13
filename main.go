package main

import (
	"fmt"
	"time"
)

func producer(q *Queue, data string) {
	msg := Message{Data: data}
	q.Send(msg)
}

func consumer(q *Queue) {
	msg := q.Receive()
	fmt.Println("Received:", msg.Data)
}

func main() {
	q := &Queue{Msgs: make(chan Message)}

	go producer(q, "Hello, world!")
	go consumer(q)

	// Give the goroutines time to finish.
	time.Sleep(1 * time.Second)
}
