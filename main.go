package main

import (
	"fmt"
	"time"
)

func producer(q *Queue, data string) {
	msg := Message{Data: data}
	if err := q.Send(msg); err != nil {
		fmt.Println("Failed to send message to queue:", err)
	}
}

func consumer(q *Queue) {
	msg, err := q.Receive()
	if err != nil {
		fmt.Println("Failed to receive message:", err)
		return
	}

	fmt.Println("Received:", msg.Data)
}

func main() {
	q := &Queue{Msgs: make(chan Message, 10)} // Added buffer to the channel

	go producer(q, "Hello from producer 1")
	go producer(q, "Hello from producer 2")
	go producer(q, "Hello from producer 3")
	go consumer(q)
	go consumer(q)
	go consumer(q)

	// Give the goroutines time to finish
	time.Sleep(1 * time.Second)

	close(q.Msgs)
}
