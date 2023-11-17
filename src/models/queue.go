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

func (q *Queue) AddNewJob(data []byte, delay int, expireAfter int) (*Job, error) {
	job, err := NewJob(data, delay, expireAfter)
	if err != nil {
		return nil, err
	}

	if job.Delay > 0 {
		q.DelayLocker.Lock()
		q.Delayed[job.Id] = job
		q.DelayLocker.Unlock()
	} else {
		q.ReadyLocker.Lock()
		q.Ready[job.Id] = job
		q.ReadyLocker.Unlock()
	}

	return job, nil
}
