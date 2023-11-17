package models

import (
	"sync"
)

type Queue struct {
	Id   int
	Name string

	DelayLocker sync.Mutex
	Delayed     map[int]*Job

	ReadyLocker sync.Mutex
	Ready       map[int]*Job

	ReservedLocker sync.Mutex
	Reserved       map[int]*Job

	IgnoredLocker sync.Mutex
	Ignored       map[int]*Job

	SubscriberLocker sync.Mutex
	Subscribers      map[int]*Subscriber
}

func NewQueue(name string) (*Queue, error) {
	queue := Queue{}
	queue.Name = name

	return &queue, nil
}
