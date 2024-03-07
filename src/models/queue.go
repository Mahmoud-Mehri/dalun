package models

import (
	"errors"
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
	queue := Queue{
		Delayed:     map[int]*Job{},
		Ready:       map[int]*Job{},
		Reserved:    map[int]*Job{},
		Ignored:     map[int]*Job{},
		Subscribers: map[int]*Subscriber{},
	}
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

func (q *Queue) DeleteJob(id int) error {
	if q.Ready[id] == nil {
		if q.Delayed[id] == nil {
			return errors.New(ERROR_JOB_NOTFOUND_MSG)
		} else {
			q.Delayed[id] = nil
			return nil
		}
	}

	q.Ready[id] = nil
	return nil
}
