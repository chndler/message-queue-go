package main

type Queue struct {
	Msgs chan Message
}

func (q *Queue) Send(msg Message) {
	q.Msgs <- msg
}

func (q *Queue) Receive() Message {
	return <-q.Msgs
}
