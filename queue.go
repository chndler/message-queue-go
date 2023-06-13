package main

import (
	"errors"
)

type Queue struct {
	Msgs chan Message
}

func (q *Queue) Send(msg Message) error {
	select {
	case q.Msgs <- msg:
		return nil
	default:
		return errors.New("queue is full")
	}
}

func (q *Queue) Receive() (Message, error) {
	select {
	case msg, ok := <-q.Msgs:
		if !ok {
			return Message{}, errors.New("queue is closed")
		}
		return msg, nil
	default:
		return Message{}, errors.New("queue is empty")
	}
}
